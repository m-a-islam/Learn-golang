package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	panic("something bad happened")
	fmt.Println("end") // this line won't print
}
