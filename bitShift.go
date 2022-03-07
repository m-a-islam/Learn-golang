package main

import (
	"fmt"
)

func main()  {
	a := 8 // 2³
	fmt.Println(a << 3) // bitshift right 2³⁺³ = 2⁶
	fmt.Println(a >> 3) // bitshift left 2³-³ = 2⁰
}