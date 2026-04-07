package main

import (
	"fmt"
	"net/http"

	"github.com/chiprek/bassurance/handlers"
)

func main() {
	fmt.Println("Initalizing Benchmark assurance")

	http.HandleFunc("/", handlers.Home)

	http.HandleFunc("/admin", handlers.Admin)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)

}
