package main

import (
	"fmt"
	"log"

	"github.com/dixonwille/wlog/v3"
	"github.com/dixonwille/wmenu/v5"
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

func enterTempreture() float64 {
	var inputTemp float64
	fmt.Print("Enter the temperature to convert: ")
	fmt.Scanln(&inputTemp)
	return inputTemp
}

func main() {
	shouldExit := false
	for !shouldExit {
		menu := wmenu.NewMenu("Choose menu item")

		menu.AddColor(wlog.Green, wlog.Yellow, wlog.Blue, wlog.Red)
		menu.ClearOnMenuRun()
		menu.LoopOnInvalid()

		menu.Option("Celsius to Fahrenheit", 1, false, nil)
		menu.Option("Fahrenheit to Celsius", 2, false, nil)
		menu.Option("Exit", 3, true, nil)

		menu.Action(func(opts []wmenu.Opt) error {
			value := opts[0].Value.(int)
			switch value {
			case 1:
				inputTemp := enterTempreture()
				convertedTemp := celsiusToFahrenheit(inputTemp)
				fmt.Printf("%.2f °C = %.2f °F\n", inputTemp, convertedTemp)
			case 2:
				inputTemp := enterTempreture()
				convertedTemp := fahrenheitToCelsius(inputTemp)
				fmt.Printf("%.2f °F = %.2f °C\n", inputTemp, convertedTemp)
			case 3:
				shouldExit = true
			}
			fmt.Scan()
			return nil
		})
		err := menu.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
