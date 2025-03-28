package main

import (
	"Backend/database"
	"Backend/models"
	"Backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	database.DB.Migrator().ColumnTypes(&models.Post{})
	database.DB.Migrator().ColumnTypes(&models.User{})
	router := gin.Default()
	routes.RegisterRoutes(router)
	port := ":8000"
	log.Println("Server running on http://localhost" + port)
	if err := router.Run(port); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
