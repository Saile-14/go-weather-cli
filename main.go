package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		LocalTime string `json:"localtime"`
	} `json:"location"`

	Current struct {
		Temperature        float64  `json:"temperature"`
		WeatherDescription []string `json:"weather_descriptions"`
	} `json:"current"`
}

func main() {
	res, err := http.Get("http://api.weatherstack.com/current?access_key=310a1ed8b542d2c6bd37eba66fa05529&query=Tokyo")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("weather api not work")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, country, current, localtime, description := weather.Location.Name, weather.Location.Country, weather.Current.Temperature, weather.Location.LocalTime, weather.Current.WeatherDescription

	message := fmt.Sprintf("Current time: %s, %s, %s: %.0fC, %s", localtime, location, country, current, description[0])

	color.Green(message)
}
