package main

import "fmt"

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func devide(a float64, b float64) float64 {
	return a / b
}

func main() {
	var a, b float64

	fmt.Print("Enter value for a: ")
	fmt.Scan(&a)
	fmt.Print("Enter value for b: ")
	fmt.Scan(&b)

	action := -1

	fmt.Print("Choose action (1 - add, 2 - subtract, 3 - multiply, 4 - divide)")
	fmt.Scan(&action)

	var result float64

	switch action {
	case 1:
		result = add(a, b)
	case 2:
		result = subtract(a, b)
	case 3:
		result = multiply(a, b)
	case 4:
		result = devide(a, b)
	}

	fmt.Println("result = ", result)
}
