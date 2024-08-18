package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to MySQL database successfully!")
}
