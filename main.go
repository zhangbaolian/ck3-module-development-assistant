package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CK3 Module Development Assistant		")
	})

	http.ListenAndServe(":8989", nil)
}
