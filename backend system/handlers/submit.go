package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func SubmitHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		content := r.FormValue("content")
		if content == "" {
			http.Error(w, "Content is required", http.StatusBadRequest)
			return
		}

		_, err := db.Exec("INSERT INTO submissions (content) VALUES (?)", content)
		if err != nil {
			http.Error(w, "Failed to insert content", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Content submitted successfully!")
	}
}
