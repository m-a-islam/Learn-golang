package main

import (
	"fmt"
)

func main() {
	sayHelloMsg("Hello", "Muhammad")
	sayByeMsg("Goodbye", "Muhammad")
}

// parameters datatypes written
func sayHelloMsg(greeting string, name string) {
	fmt.Println(greeting, name)
}

func sayByeMsg(msg, name string) {
	fmt.Println(msg, name)
}
