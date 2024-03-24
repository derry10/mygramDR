package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartDB() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.UserRegister)
		userGroup.POST("/login", controllers.UserLogin)
		userGroup.PUT("/:userId", middlewares.Authentication(), controllers.UserUpdate)
		userGroup.DELETE("/delete", middlewares.Authentication(), controllers.UserDelete)
	}

	photoGroup := r.Group("/photos")
	{
		photoGroup.Use(middlewares.Authentication())
		photoGroup.POST("/", controllers.CreatePhoto)
		photoGroup.GET("/", controllers.GetPhoto)
		photoGroup.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoGroup.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentGroup := r.Group("/comments")
	{
		commentGroup.Use(middlewares.Authentication())
		commentGroup.POST("/", controllers.CreateComment)
		commentGroup.GET("/", controllers.GetComment)
		commentGroup.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentGroup.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialmediasGroup := r.Group("/socialmedias")
	{
		socialmediasGroup.Use(middlewares.Authentication())
		socialmediasGroup.POST("/", controllers.CreateSocialMedia)
		socialmediasGroup.GET("/", controllers.GetSocialMedia)
		socialmediasGroup.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediasGroup.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
