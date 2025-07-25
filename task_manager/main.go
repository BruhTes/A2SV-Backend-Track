package main

import (
	"log"
	"task_manager/data"
	"task_manager/router"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Connect to MongoDB
	data.InitMongoConnection()
	data.InitUserCollection(&mongo.Database{})

	r := router.SetupRouter()
	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
