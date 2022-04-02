package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}
	// defer is always execute at the end of the runtime.
	// it seems it is closing first and opening the resources but it is not.
	// it will defer this line at the end of the runtime of this program
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
