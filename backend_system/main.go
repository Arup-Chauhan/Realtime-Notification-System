package main

import (
	"Realtime-Notification-System/backend_system/database"
	"Realtime-Notification-System/backend_system/server"
)

func main() {
	database.InitDB()
	server.Router()

}
