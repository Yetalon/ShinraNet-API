package routes

import (
	"Backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("", controllers.CreateUser)
		users.GET("/:id", controllers.GetUserByID)
		users.GET("", controllers.GetUserByName)
		users.PUT("/:id", controllers.UpdateUser)
		users.PATCH("/:id", controllers.PatchUserName)
		users.PATCH("/:id/is_admin", controllers.PatchAdmin)
		users.PATCH("/:id/image_url", controllers.PatchImage)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	media := router.Group("/posts")
	{
		media.POST("", controllers.CreatePost)
		media.GET("/:id", controllers.GetPostById)
		media.GET("", controllers.GetPostsByTitle)
		media.GET("/user/:id", controllers.GetPostByUserId)
		media.PUT("/:id", controllers.UpdatePost)
		media.PATCH("/:id/title", controllers.PatchPostTitle)
		media.PATCH("/:id/post_text", controllers.PatchPostText)
		media.PATCH("/:id/increment_likes")
		media.PATCH("/:id/decrement_likes")
		media.DELETE("/:id", controllers.DeletePost)
	}
}
