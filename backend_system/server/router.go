package server

import (
	"Realtime-Notification-System/backend_system/handlers"
	middleware "Realtime-Notification-System/middleware_layer"
	"database/sql"
	"net/http"
)

func Router(db *sql.DB) error {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Set up the routes with middleware
	mux.Handle("/submit", middleware.EnableCORS(handlers.SubmitHandler(db)))

	// Start the server using the ServeMux
	return http.ListenAndServe(":8080", mux)
}
