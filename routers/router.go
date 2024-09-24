package routers

import (
	"final-project-golang/controllers"
	"final-project-golang/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	//URL untuk manage user
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	//URL untuk manage social media
	socialmediaRouter := r.Group("/socialmedia")
	{
		socialmediaRouter.Use(middleware.Authentication())
		socialmediaRouter.POST("/create", controllers.CreateSocialMedia)
		socialmediaRouter.GET("/", controllers.GetSocialMedia)
		socialmediaRouter.GET("/:socialmediaId", controllers.GetIdSocialMedia)
		socialmediaRouter.PUT("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediaRouter.DELETE("/:socialMediaId", middleware.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	//URL untuk manage photo
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/create", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.GET("/:photoId", controllers.GetByIdPhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controllers.DeletePhoto)
	}

	//URL untuk manage comment
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/create", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.GET("/:commentId", controllers.GetByIdComment)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controllers.DeleteComment)
	}

	return r
}
