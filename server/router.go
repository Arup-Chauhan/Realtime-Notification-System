package server

import (
	"fmt"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Webserver: Level 1")
	})

	fmt.Println("Server running on standard 8080 localhost port")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
