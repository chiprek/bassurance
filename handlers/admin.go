package handlers

import (
	"fmt"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Benchmark assurance admin menu")
}
