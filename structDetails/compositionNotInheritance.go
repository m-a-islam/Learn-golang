package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

// in go it doesn't have inheritence but it has composition
type Bird struct { // it is Has relationship
	Animal   // taking the Animal struct
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

}
