package routes

import (
	"Backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/users/", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUserByID)
	router.GET("/users", controllers.GetUserByName)
}
