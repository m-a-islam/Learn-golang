package main

import (
	"fmt"
)

func main() {
	var roman string
	roman = "MCMXCIV"
	toDeci(roman)
	/*for _, letter := range roman {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
		if strings.ContainsAny(string(letter), "L | I | V | X | C | D | M") {
			if string(letter) == "I" {
				res += 1
			}
			if string(letter) == "V" {
				res += 5
			}
			if string(letter) == "X" {
				res += 10
			}
			if string(letter) == "L" {
				res += 50
			}
			if string(letter) == "C" {
				res += 100
			}
			if string(letter) == "D" {
				res += 500
			}
			if string(letter) == "M" {
				res += 1000
			}
		}
	}
	fmt.Println(res)
	*/
}

func toDeci(s string) {

	roman := map[string]int{ // [typeOfKey]typeOfValue
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0
	for key, v := range s {
		if roman[string(v)] < roman[string(s[key+1])] {
			res += roman[string(s[key+1])] - roman[string(v)]
		} else {
			res += roman[string(v)]
		}
	}
	fmt.Println(res)
}
