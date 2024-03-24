# Weather Station Distributor

This project is a weather station distributor that allows users to collect and distribute weather data from a weather station.
The current state only sends the data to a windguru.cz account.

## Example request from a weather station
```
http://localhost:3000/weatherstation/updateweatherstation.php?ID=idstring&PASSWORD=password&action=updateraww&realtime=1&rtfreq=5&dateutc=now&baromin=29.92&tempf=43.5&dewptf=40.8&humidity=90&windspeedmph=3.8&windgustmph=4.2&winddir=234&rainin=0.0&dailyrainin=0.0&solarradiation=0.0&UV=0.0&indoortempf=71.9&indoorhumidity=55
```
