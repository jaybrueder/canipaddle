package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"canipaddle.com/internal/models"
)

func HomeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Venzone Tagliamento measure API endpoint
		url := "https://monitor.protezionecivile.fvg.it/api/stations/526/measures/latest"
		actualMeasure := 0.0
		desiredMeasure := 0.4

		// Make the GET request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making GET request:", err)
			return fmt.Errorf("error making GET request: %w", err)
		}
		defer resp.Body.Close()

		// Check if the request was successful
		if resp.StatusCode != http.StatusOK {
			fmt.Println("API request failed with status code:", resp.StatusCode)
			return fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
		}

		// Create a variable of our APIResponse type
		var apiResp models.APIResponse

		// Decode the JSON response into struct
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&apiResp); err != nil {
			fmt.Println("Error decoding JSON response:", err)
			return fmt.Errorf("error decoding JSON response: %w", err)
		}

		fmt.Printf("API Result: %s\n", apiResp.Result)
		fmt.Printf("Found %d measures\n", len(apiResp.Measures))

		if len(apiResp.Measures) > 0 {
			actualMeasure = apiResp.Measures[2].Value
		}

		canIPaddle := false

		if actualMeasure > desiredMeasure {
			canIPaddle = true
		}

		// Render the template with the monsters
		component := Home(canIPaddle, actualMeasure, desiredMeasure)
		return component.Render(c.Request().Context(), c.Response().Writer)
	}
}
