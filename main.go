package main

import (
	"fmt"
	"log"
	"net/http"

	"htmx/handlers"
)

func main() {
	fmt.Println("Go app...")

	// define handlers
	http.HandleFunc("/", handlers.H1)
	http.HandleFunc("/add-film/", handlers.H2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
