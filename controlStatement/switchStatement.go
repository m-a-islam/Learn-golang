package main

import (
	"fmt"
)

func main() {
	switch 2 {
	case 1, 3, 5:
		fmt.Println("odd")
	case 2, 4, 6:
		fmt.Println("even")
	default:
		fmt.Println("other number")
	}
}
