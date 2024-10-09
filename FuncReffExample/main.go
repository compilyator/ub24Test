package main

import "fmt"

func main() {
	var a, b float64

	fmt.Print("Enter value for a: ")
	fmt.Scan(&a)
	fmt.Print("Enter value for b: ")
	fmt.Scan(&b)

	action := -1

	fmt.Print("Choose action (1 - add, 2 - subtract, 3 - multiply, 4 - divide)")
	fmt.Scan(&action)

	functions := map[int](func(float64, float64) float64){
		1: func(f1, f2 float64) float64 { return f1 + f2 },
		2: func(f1, f2 float64) float64 { return f1 - f2 },
		3: func(f1, f2 float64) float64 { return f1 * f2 },
		4: func(f1, f2 float64) float64 { return f1 / f2 },
	}

	result := functions[action](a, b)

	fmt.Println("result = ", result)
}
