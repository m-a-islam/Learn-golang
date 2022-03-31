package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}

	for key, value := range s {
		fmt.Println(key, value)
	}

	statePopulations := map[string]int{ // [typeOfKey]typeOfValue
		"Dresden":       234344,
		"Berlin":        745894,
		"FrankfurtMain": 37485,
		"NRW":           98456,
	}
	for key, value := range statePopulations {
		fmt.Println(key, value)
	}

	// loop through the string

	sentence := "hello this is Muhammad"

	for key, value := range sentence {
		fmt.Println(key, string(value))
	}
}
