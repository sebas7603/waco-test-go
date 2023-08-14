package cmd

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/config"
	"github.com/sebas7603/waco-test-go/pkg/controllers"
)

var err error

func Start() error {
	err = config.InitialConfig()
	if err != nil {
		fmt.Println("Error in initial config", err)
		return err
	}

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)

		userGroup := api.Group("/users")
		{
			userGroup.GET("/", controllers.IndexUsers)
			userGroup.GET("/:id", controllers.ShowUser)
		}

		characterGroup := api.Group("/characters")
		{
			characterGroup.GET("/", controllers.IndexCharacters)
			characterGroup.GET("/:id", controllers.ShowCharacter)
		}
	}

	router.Run(":" + os.Getenv("APP_PORT"))
	return nil
}
