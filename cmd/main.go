package cmd

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/config"
	"github.com/sebas7603/waco-test-go/routes"
)

var err error

func Start() error {
	err = config.InitialConfig()
	if err != nil {
		fmt.Println("Error in initial config", err)
		return err
	}

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run("0.0.0.0:" + os.Getenv("PORT"))
	return nil
}
