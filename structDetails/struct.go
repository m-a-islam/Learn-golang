package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John Pattinson",
		companions: []string{
			"Kathleen",
			"Saad",
			"Micha",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[2])
}
