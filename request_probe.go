package main

type RequestProbe struct {
	ID             string  `json:"ID" query:"ID"`
	Password       string  `json:"password" query:"PASSWORD"`
	Temperature    float64 `json:"temperature" query:"tempf"`
	DewPoint       float64 `json:"dew_point" query:"dewptf"`
	Humidity       float64 `json:"humidity" query:"humidity"`
	WindSpeed      float64 `json:"wind_speed" query:"windspeedmph"`
	WindGust       float64 `json:"wind_gust" query:"windgustmph"`
	WindDirection  float64 `json:"wind_direction" query:"winddir"`
	Pressure       float64 `json:"pressure" query:"baromin"`
	Rain           float64 `json:"rain" query:"rainin"`
	DailyRain      float64 `json:"daily_rain" query:"dailyrainin"`
	SolarRadiation float64 `json:"solarradiation" query:"solarradiation"`
	UV             float64 `json:"uv" query:"UV"`
}

func (r *RequestProbe) fahrenheitToCelcius() {
	r.Temperature = (r.Temperature - 32) * 5 / 9
	r.DewPoint = (r.DewPoint - 32) * 5 / 9
}

func (r *RequestProbe) mphToKnots() {
	r.WindSpeed = r.WindSpeed * 0.868976
	r.WindGust = r.WindGust * 0.868976
}
