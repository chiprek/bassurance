package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chiprek/bassurance/handlers"
)

func main() {
	fmt.Println("Initalizing Benchmark assurance")

	db, err := initDB()
	if err != nil {
		log.Fatalf("could not initalize database: %v", err)
	}
	defer db.Close()

	app := &handlers.App{
		DB: db,
	}

	http.HandleFunc("/", app.Home)

	//http.HandleFunc("/admin", app.Admin)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
