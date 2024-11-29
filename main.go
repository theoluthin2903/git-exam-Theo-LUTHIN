package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "route: %q", r.URL.Path)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	// to add : color functionality
	log.Fatal(http.ListenAndServe(":8080", nil))
}
