package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	salt := os.Getenv("SALT")
    uid := os.Getenv("UID")
    password := os.Getenv("PASSWORD")
	stationId := os.Getenv("STATION_ID")
	stationPassword := os.Getenv("STATION_PASSWORD")

	app := fiber.New()

	app.Get("/weatherstation/updateweatherstation.php", func(c *fiber.Ctx) error {
		probe := &RequestProbe{}

		if err := c.QueryParser(probe); err != nil {
			return err
		}

		if probe.ID != stationId || probe.Password != stationPassword {
			return c.SendStatus(http.StatusUnauthorized)
		}

		probe.fahrenheitToCelcius()
		probe.mphToKnots()

		baseUrl, _  := url.Parse("http://www.windguru.cz/upload/api.php")
		params := url.Values{}
		params.Add("uid", uid)
		params.Add("salt", salt)
		params.Add("hash", getHash(salt, uid, password))
		params.Add("wind_avg", fmt.Sprintf("%f", probe.WindSpeed))
		params.Add("wind_dir", fmt.Sprintf("%f", probe.WindDirection))
		params.Add("temp", fmt.Sprintf("%f", probe.Temperature))

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
