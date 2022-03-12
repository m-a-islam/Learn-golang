package main

import (
	"fmt"
)

func main() {
	// declaring struct anonymously // initializing the struct
	aDoctor := struct{ name string }{name: "John pattinson"}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Steve Jones"
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}
