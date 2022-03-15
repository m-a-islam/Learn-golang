package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := -10

	if returnTrue() || guess < 1 || guess > 100 { // go doesnt check for second operator in logical OR
		fmt.Println("The guess must be between 1 and 100")
	}
	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You got it")
		}
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
