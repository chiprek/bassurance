package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Initalizing Benchmark assurance")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Benchmark assurance is running")
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)

}
