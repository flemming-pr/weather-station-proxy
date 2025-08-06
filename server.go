package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	salt := os.Getenv("SALT")
	uid := os.Getenv("UID")
	password := os.Getenv("PASSWORD")
	stationId := os.Getenv("STATION_ID")
	stationPassword := os.Getenv("STATION_PASSWORD")

	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) error {
		filename := "current.json"

		jsonData, err := os.ReadFile(filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error reading file: %v", err))
		}

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

		return c.Send(jsonData)
	})

	// app.Get("/*", func(c *fiber.Ctx) error {
	// 	log.Printf("Request received: %s %s", c.Method(), c.Path())
	// 	log.Printf("Query parameters: %v", c.AllParams())
	// 	return c.SendStatus(fiber.StatusOK)
	// })

	app.Get("/weatherstation/updateweatherstation.php", func(c *fiber.Ctx) error {
		probe := &RequestProbe{}

		log.Printf("ping from station")

		if err := c.QueryParser(probe); err != nil {
			return err
		}

		log.Printf(probe.ID)
		if probe.ID != stationId || probe.Password != stationPassword {
			return c.SendStatus(http.StatusUnauthorized)
		}

		probe.mphToKnots()
		probe.fahrenheitToCelcius()

		writeToFile(*probe)

		baseUrl, _ := url.Parse("http://www.windguru.cz/upload/api.php")
		params := url.Values{}
		params.Add("uid", uid)
		params.Add("salt", salt)
		params.Add("hash", getHash(salt, uid, password))
		params.Add("wind_avg", fmt.Sprintf("%f", probe.WindSpeed))
		params.Add("wind_max", fmt.Sprintf("%f", probe.WindGust))
		params.Add("wind_direction", fmt.Sprintf("%f", probe.WindDirection))
		params.Add("temperature", fmt.Sprintf("%f", probe.Temperature))
		params.Add("rh", fmt.Sprintf("%f", probe.Humidity))
		params.Add("mslp", fmt.Sprintf("%f", probe.Pressure))

		baseUrl.RawQuery = params.Encode()

		resp, err := http.Get(baseUrl.String())

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != 200 {
			fmt.Println(string(body))
		}

		return c.Send(body)
	})

	log.Fatal(app.Listen(":3000"))
}

func getHash(salt string, uid string, password string) string {
	data := []byte(salt + uid + password)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func writeToFile(probe RequestProbe) {
	probe.hidePassword()
	jsonData, err := json.Marshal(probe)
	if err != nil {
		log.Fatalf("Error encoding %v", err)
	}

	filename := "current.json"

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
