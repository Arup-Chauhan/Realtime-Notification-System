package server

import (
	"Realtime-Notification-System/backend-system/database"
	"Realtime-Notification-System/backend-system/handlers"

	"fmt"
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Webserver: Level 1")
	})

	http.HandleFunc("/submit", handlers.SubmitHandler(database.DB))

	fmt.Println("Server running on standard 8080 localhost port")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
