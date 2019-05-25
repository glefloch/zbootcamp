package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting zBootcamp server...")

	http.HandleFunc("/", ZHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func ZHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling zRequest..")
	fmt.Fprintf(w, `{"hello": "zenika"}`)
}
