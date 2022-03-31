package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	k := 0
	for ; k < 5; k++ {
		fmt.Println(k)
	}

	// 2D loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// special ForLoop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	l := 0
	for {
		fmt.Println(l)
		l++
		if l == 5 {
			break
		}
	}
}
