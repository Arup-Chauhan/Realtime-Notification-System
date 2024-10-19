package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Realtime-Notification-System/backend_system/handlers"
	middleware "Realtime-Notification-System/middleware_layer"
)

func Router(db *sql.DB) error {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Set up routes with middleware
	mux.Handle("/submit", middleware.EnableCORS(handlers.SubmitHandler(db)))

	// Create a custom server instance
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Run the server in a goroutine to avoid blocking
	go func() {
		log.Println("Server running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for termination signal
	<-stop
	log.Println("Shutting down server...")

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
