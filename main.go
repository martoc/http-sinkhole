package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGone)
}

func main() {
	fmt.Printf("Starting http-sinkhole at 8080 port")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}