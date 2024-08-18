package main

import (
	"Realtime-Notification-System/backend-system/database"
	"Realtime-Notification-System/backend-system/server"
)

func main() {
	database.InitDB()
	server.Router()

}
