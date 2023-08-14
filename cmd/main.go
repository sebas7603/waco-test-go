package cmd

import (
	"fmt"
	"os"

	"github.com/sebas7603/waco-test-go/config"
)

var err error

func Start() error {
	err = config.InitialConfig()
	if err != nil {
		fmt.Println("Error in initial config", err)
		return err
	}

	fmt.Println("Rick And Morty API:", os.Getenv("RM_API_URL"))
	return nil
}
