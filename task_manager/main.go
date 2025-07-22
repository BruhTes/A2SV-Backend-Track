package main

import (
	"log"
	"task_manager/data"
	"task_manager/router"
)

func main() {
	// Connect to MongoDB
	data.InitMongoConnection()

	r := router.SetupRouter()
	log.Println("🚀 Server is running on http://localhost:8080")
	r.Run(":8080")
}
