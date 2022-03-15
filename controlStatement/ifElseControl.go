package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("This test is true")
	}
	statePopulations := make(map[string]int) // [typeOfKey]typeOfValue, Here make() is just initializing the associative array
	statePopulations = map[string]int{       // [typeOfKey]typeOfValue
		"Dresden":       234344,
		"Berlin":        745894,
		"FrankfurtMain": 37485,
		"NRW":           98456,
	}
	//population, ok := statePopulations["FrankfurtMain"]
	if population, ok := statePopulations["FrankfurtMain"]; ok {
		fmt.Println(population)
	}
}
