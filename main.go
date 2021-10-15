package main

import (
	"test-v2/models"
	"test-v2/routes"

)

func main() {
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)

	r.Run()
}

