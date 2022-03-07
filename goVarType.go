package main

import "fmt"

// this upper case I is exposed to other package
var I int = 100

// this j is global but within this package
var j int = 20

func main() {
	// this i is local variable and not exposed
	var i float64
	i = 40.0
	j = I + j

	fmt.Printf("%v %T\n", i, i)
}
