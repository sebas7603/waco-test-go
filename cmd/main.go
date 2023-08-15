package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/config"
	"github.com/sebas7603/waco-test-go/pkg/controllers"
	"github.com/sebas7603/waco-test-go/pkg/middlewares"
)

var err error

func Start() error {
	err = config.InitialConfig()
	if err != nil {
		fmt.Println("Error in initial config", err)
		return err
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"api": "waco-test",
		})
	})
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// userGroup := api.Group("/users")
		// {
		// 	userGroup.GET("/", controllers.IndexUsers)
		// 	userGroup.GET("/:id", controllers.ShowUser)
		// }

		privateGroup := api.Group("/")
		{
			privateGroup.Use(middlewares.AuthMiddleware())
		}

		privateGroup.POST("/renew-token", controllers.RenewToken)

		profileGroup := privateGroup.Group("/profile")
		{
			profileGroup.GET("/", controllers.ShowProfile)
			profileGroup.PUT("/", controllers.UpdateProfile)
			profileGroup.POST("/favorite", controllers.AddFavoriteByID)
		}

		characterGroup := privateGroup.Group("/characters")
		{
			characterGroup.GET("/", controllers.IndexCharacters)
			characterGroup.GET("/:id", controllers.ShowCharacter)
		}
	}

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
	return nil
}
