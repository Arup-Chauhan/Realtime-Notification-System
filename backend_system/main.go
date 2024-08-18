package main

import (
	"Realtime-Notification-System/database"
	"Realtime-Notification-System/server"
)

func main() {
	database.InitDB()
	server.Router()

}
