package main

import (
	"Realtime-Notification-System/backend_system/database"
	"Realtime-Notification-System/backend_system/server"
	"log"
)

func main() {
	// Initialize the database connection
	db := database.InitDB()

	// Start the server using the custom router
	log.Println("Server starting on port 8080...")
	log.Fatal(server.Router(db))
}
