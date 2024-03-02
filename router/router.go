package router

import (
	"profile-picture-api/controllers"
	"profile-picture-api/database"
	"profile-picture-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Profile Picture API!",
			"author":  "Fikri Naufal Hamdi",
		})
	})

	database.ConnectDatabase()

	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.Login)
	router.POST("/users/logout", controllers.Logout)

	authorized := router.Group("/")
	authorized.Use(middlewares.CheckAuth())
	{
		authorized.GET("/users/:id", controllers.GetUserbyID)
		authorized.POST("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)

		authorized.GET("/photos", controllers.GetAllPhotos)
		authorized.POST("/photos", controllers.CreatePhoto)
		authorized.PUT("/photos/:id", controllers.UpdatePhoto)
		authorized.DELETE("/photos/:id", controllers.DeletePhoto)
	}

	return router
}
