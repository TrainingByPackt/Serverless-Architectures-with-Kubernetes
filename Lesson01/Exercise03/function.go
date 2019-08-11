package main

import (
	"fmt"
	"net/http"
)

func WelcomeServerless(w http.ResponseWriter, r *http.Request) {

	names, ok := r.URL.Query()["name"]
    
    if ok && len(names[0]) > 0 {
        fmt.Fprintf(w, names[0] + ", Hello Serverless World!")
	} else {
		fmt.Fprintf(w, "Hello Serverless World!")
	}
}