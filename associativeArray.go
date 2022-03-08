package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int) // [typeOfKey]typeOfValue, Here make() is just initializing the associative array
	statePopulations = map[string]int{ // [typeOfKey]typeOfValue
		"Dresden": 234344,
		"Berlin": 745894,
		"FrankfurtMain": 37485,
		"NRW": 98456,
	}
	statePopulations["Fulda"] = 777458 // inserting new elements in the associative array
	delete(statePopulations, "FrankfurtMain") // deleting from array
	population, ok := statePopulations["FrankfurtMain"] // check if the array key isset in the array, if not it will print 0, false
	fmt.Println(population, ok)
}