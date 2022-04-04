package main

import "fmt"

func main() {
	greeting := "hello"
	name := "Muhammad"
	sayGreetings(&greeting, &name)

	fmt.Println("--main funtion--")
	fmt.Println(name)
}

func sayGreetings(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Aurpi"
	fmt.Println("--from funtion--")
	fmt.Println(*name)

}
