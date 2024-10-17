package main

import (
	"fmt"
)

const (
	celsiusToFahrenheitFactor = 9.0 / 5.0
	fahrenheitToCelsiusFactor = 5.0 / 9.0
	offset                    = 32.0
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * celsiusToFahrenheitFactor) + offset
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - offset) * fahrenheitToCelsiusFactor
}

func main() {
	for {
		fmt.Println("Temperature Converter")
		fmt.Println("1. Celsius to Fahrenheit")
		fmt.Println("2. Fahrenheit to Celsius")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option (1-3): ")

		var option int
		fmt.Scanln(&option)

		if option == 3 {
			fmt.Println("Exiting the program.")
			break
		}

		var inputTemp float64
		fmt.Print("Enter the temperature to convert: ")
		fmt.Scanln(&inputTemp)

		switch option {
		case 1:
			convertedTemp := celsiusToFahrenheit(inputTemp)
			fmt.Printf("%.2f °C = %.2f °F\n", inputTemp, convertedTemp)
		case 2:
			convertedTemp := fahrenheitToCelsius(inputTemp)
			fmt.Printf("%.2f °F = %.2f °C\n", inputTemp, convertedTemp)
		default:
			fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		}

		fmt.Println()
	}
}
