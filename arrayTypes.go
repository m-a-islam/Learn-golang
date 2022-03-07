package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}
	gradesHoldAny := [...]int{97, 85, 93, 88}

	fmt.Printf("Grades: %v \n", grades)
	fmt.Printf("Grades: %v \n", gradesHoldAny)
	fmt.Printf("length Grades: %v \n", len(gradesHoldAny))


	a := make([]int, 3, 100)

	fmt.Println(a)
	fmt.Printf("Length: %v \n", len(a))
	fmt.Printf("Capacity: %v \n", cap(a))

}
