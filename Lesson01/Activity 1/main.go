package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the \xF0\x9F\x9A\xB2 finder..")
	http.HandleFunc("/", FindBikes)
	fmt.Println("Function handlers are registered.")

	http.ListenAndServe(":8080", nil)
}
