package main

import (
	"Backend/database"
	"Backend/models"
	"Backend/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Post{})
	router := gin.Default()
	routes.RegisterRoutes(router)
	port := ":8000"
	log.Println("Server running on http://localhost" + port)
	if err := router.Run(port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
