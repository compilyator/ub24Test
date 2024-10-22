package main

import "fmt"

func main() {
	fmt.Println("Welcome to guess game. Choise number between 0 and 100 and memorize it")
	low := 0
	higher := 100
	attempts := 0

	for {
		guess := (low + higher) / 2
		attempts++
		fmt.Println("Attempt number ", attempts, ". Is you number ", guess, "?")
		fmt.Println("Enter your answer:")
		fmt.Println("- e for equal")
		fmt.Println("- h if your number is higher than mine")
		fmt.Println("- l if your number lower than mine)")
		var input string
		fmt.Scan(&input)
		switch input {
		case "e":
			fmt.Println("I won with in ", attempts, " attempts")
			return
		case "h":
			low = guess
		case "l":
			higher = guess
		}
	}
}
