package main

import (
	"Library-Management/controllers"
	"Library-Management/services"
)

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)
	controller.Start()
}
