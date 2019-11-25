package p
import (
    "fmt"
    "net/http"
)
func HelloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World from Google Cloud Functions!")
    return
}