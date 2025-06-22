package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go webhook!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started on :10000")
	log.Fatal(http.ListenAndServe(":10000", nil))
}
