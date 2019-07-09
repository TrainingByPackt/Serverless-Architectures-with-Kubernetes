package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the 🚲 finder..")
	http.HandleFunc("/", FindBikes)
	fmt.Println("Function handlers are registered.")

	http.ListenAndServe(":8080", nil)
}
