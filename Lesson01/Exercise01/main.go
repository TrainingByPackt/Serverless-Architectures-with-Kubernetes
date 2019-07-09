package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the serverless environment..")
	http.HandleFunc("/", WelcomeServerless)
	fmt.Println("Function handlers are registered.")

	http.ListenAndServe(":8080", nil)
}
