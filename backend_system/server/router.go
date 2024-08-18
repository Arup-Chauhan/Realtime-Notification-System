package server

import (
	"fmt"
	"log"
	"net/http"

	"Realtime-Notification-System/backend_system/database"
	"Realtime-Notification-System/backend_system/handlers"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Webserver: Level 1")
	})

	http.HandleFunc("/submit", handlers.SubmitHandler(database.DB))

	fmt.Println("Server running on standard 8080 localhost port")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
