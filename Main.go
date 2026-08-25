package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Weather struct {
	airTemp, airPressure, precipitation, windSpeed              string
	windDirection, humidity, dewPoint, soilMoisture, cloudCover string
}

func main() {

	WeatherData := Weather{
		airTemp:       "NULL",
		airPressure:   "NULL",
		precipitation: "NULL",
		windSpeed:     "NULL",
		windDirection: "NULL",
		humidity:      "NULL",
		dewPoint:      "NULL",
		soilMoisture:  "NULL",
		cloudCover:    "NULL",
	}

	fmt.Println("--- Weather Station ---")

	scanner := bufio.NewScanner(os.Stdin) // Read input from Terminal

	for scanner.Scan() { //keep taking input

		input := scanner.Text()

		if input == "get" {
			fmt.Println("airTemp:" + WeatherData.airTemp)
			fmt.Println("airPressure:" + WeatherData.airPressure)
			fmt.Println("precipitation:" + WeatherData.precipitation)
			fmt.Println("windSpeed:" + WeatherData.windSpeed)
			fmt.Println("windDirection:" + WeatherData.windDirection)
			fmt.Println("humidity:" + WeatherData.humidity)
			fmt.Println("dewPoint:" + WeatherData.dewPoint)
			fmt.Println("soilMoisture:" + WeatherData.soilMoisture)
			fmt.Println("cloudCover:" + WeatherData.cloudCover)

		} else if input == "clear" {
			WeatherData.airTemp = "NULL"
			WeatherData.airPressure = "NULL"
			WeatherData.precipitation = "NULL"
			WeatherData.windSpeed = "NULL"
			WeatherData.windDirection = "NULL"
			WeatherData.humidity = "NULL"
			WeatherData.dewPoint = "NULL"
			WeatherData.soilMoisture = "NULL"
			WeatherData.cloudCover = "NULL"

		} else if input == "exit" {
			fmt.Println("Exiting...")
			break
		} else {
			x := strings.Split(input, ",")

			if len(x) != 2 {
				continue
			}

			id, err := strconv.Atoi(x[0])

			if err != nil {
				continue
			}

			value := x[1]

			switch id {
			case 1:
				WeatherData.airTemp = value
			case 2:
				WeatherData.airPressure = value
			case 7:
				WeatherData.precipitation = value
			case 11:
				WeatherData.windSpeed = value
			case 12:
				WeatherData.windDirection = value
			case 13:
				WeatherData.humidity = value
			case 14:
				WeatherData.dewPoint = value
			case 15:
				WeatherData.soilMoisture = value
			case 22:
				WeatherData.cloudCover = value

			}
		}

	}
}
