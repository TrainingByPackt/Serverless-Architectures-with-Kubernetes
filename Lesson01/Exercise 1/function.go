package main

import (
	"fmt"
	"net/http"
)

func WelcomeServerless(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Serverless World!")
}