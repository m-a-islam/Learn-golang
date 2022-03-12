package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	// get the validation text and we can further do our task by reading the Tag
	fmt.Println(field.Tag)

}
