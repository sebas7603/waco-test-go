package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/pkg/controllers"
	"github.com/sebas7603/waco-test-go/pkg/middlewares"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"api": "waco-test",
		})
	})

	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		privateGroup := api.Group("/")
		privateGroup.Use(middlewares.AuthMiddleware())
		{
			// userGroup := api.Group("/users")
			// {
			// 	userGroup.GET("/", controllers.IndexUsers)
			// 	userGroup.GET("/:id", controllers.ShowUser)
			// }

			privateGroup.POST("/renew-token", controllers.RenewToken)
			privateGroup.POST("/change-password", controllers.ChangePassword)

			profileGroup := privateGroup.Group("/profile")
			{
				profileGroup.GET("/", controllers.ShowProfile)
				profileGroup.PUT("/", controllers.UpdateProfile)
				profileGroup.POST("/favorite", controllers.AddFavoriteByID)
				profileGroup.DELETE("/favorite", controllers.RemoveFavoriteByID)
			}

			characterGroup := privateGroup.Group("/characters")
			{
				characterGroup.GET("/", controllers.IndexCharacters)
				characterGroup.GET("/:id", controllers.ShowCharacter)
			}
		}
	}
}
