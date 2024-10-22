package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	numberToGuess := randGenerator.Intn(100)

	fmt.Println("Welcome to guess number game")
	fmt.Println("I have picked a number between 1 and 100. You should guess it")
	var userGuess, attempts int
	for {
		fmt.Print("Enter you guess: ")
		_, err := fmt.Scanf("%d", &userGuess)
		if err != nil {
			fmt.Println("Wrong input. Try again")
			continue
		}
		attempts++

		if userGuess == numberToGuess {
			fmt.Printf("You win with in %d attempts", attempts)
			break
		} else if userGuess < numberToGuess {
			fmt.Printf("Attemp: %d. You number lower than imagened", attempts)
		} else if userGuess > numberToGuess {
			fmt.Printf("Attemp: %d. You number greater than imagened", attempts)
		}
	}
}
