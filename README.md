# Weather Forecast CLI

A simple and efficient command-line weather forecasting tool written in Go. It fetches current weather conditions and forecasts for any given city using the WeatherAPI. The application highlights the hourly forecast, especially focusing on rain chances, to help you plan your day better.

## Features

- Fetch current weather conditions including temperature and general weather status.
- Get hourly weather forecasts including temperature and chance of rain.
- Highlights in red the hourly forecasts with a high chance of rain for easy visibility.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of Go.
- You have a basic understanding of command-line tools and Go programming language.
- You have obtained an API key from WeatherAPI and replace `"your_api_key_here"` in the source code with your actual API key.

## Installing Weather Forecast CLI

To install the Weather Forecast CLI, follow these steps:

Linux and macOS:

```bash
git clone https://github.com/yourgithubusername/weather-forecast-cli.git
cd weather-forecast-cli
go build

./weather-forecast-cli [city name]
```
