# Weather CLI

A simple **Go** command-line application that fetches and displays the current weather for a specified location using the [Weatherstack API](https://weatherstack.com/).

## Features

- Fetches current weather conditions (temperature, description).
- Displays the local time at the queried location.
- If no location is provided, **Tokyo** is used by default 💯.
- Colored console output for fun :)

## Requirements

1. [Go](https://golang.org/dl/) installed on your system (1.16+ recommended).
2. A [Weatherstack API key](https://weatherstack.com/) (free tier available).
3. A `.env` file containing your API key, or an environment variable.

## Setup

1. **Clone or download** this repository.

2. **Install dependencies**:  
   ```bash
   go get github.com/fatih/color
   go get github.com/joho/godotenv
   
3. **Create a .env file at the root of your project with the following content:**

  WEATHER_API_KEY=your_weatherstack_api_key_here

## Usage
By running go run ./main.go, it will by default give you the weather, current time and current temperature of Tokyo.
If you run go run ./main.go <Location> , you can check your own location 🔥 😄

enjoyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy
