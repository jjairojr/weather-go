package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
		Tz_id   string  `json:"tz_id"`
	} `json:"location"`

	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch    int     `json:"time_epoch"`
				TempC        float64 `json:"temp_c"`
				ChanceOfRain float64 `json:"chance_of_rain"`
				Conditon     struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	}
}

func main() {
	city := "Goiania"

	if len(os.Args) > 1 {
		city = os.Args[1]
	}

	req, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=54988f197d3f4349b4301127241104&q=" + city)

	if err != nil {
		log.Fatal("Error fetching weather data")
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)

	if err != nil {
		log.Fatal("Error reading weather data")
	}

	var weather Weather

	err = json.Unmarshal(body, &weather)

	if weather.Location.Name == "" {
		log.Fatal("City not found")
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf("%s, %s, %.0fC %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(int64(hour.TimeEpoch), 0)

		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("%s - %.0fC, %.0f, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Conditon.Text,
		)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
