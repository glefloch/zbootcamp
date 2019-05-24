package main

import (
	"fmt"
	"net/http"
)

func main() {

	content := "Zenika"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s", content)
	})
	http.ListenAndServe(":3000", nil)
}
