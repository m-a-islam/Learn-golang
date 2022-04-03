package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We are on Go Web development page"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil { // if we see the port is assigned by others then the program will panic
		panic(err.Error())
	}
}
