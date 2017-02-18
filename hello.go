package main

import (
	"fmt"
	"net/http"
)

const VERSION = "1.0.0"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! %s", VERSION)
}

func main() {
	fmt.Print("Go to http://localhost:3000/\n")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}
