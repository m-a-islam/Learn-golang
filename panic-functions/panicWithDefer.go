// the order of the function is Defer will execute before panic

package main

import (
	"fmt"
)

func main() {
	fmt.Println("begin")
	defer fmt.Println("this is the defer function")
	panic("Something bad happened!!!!")

	fmt.Println("Stop")
}
