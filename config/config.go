package config

import (
	"github.com/joho/godotenv"
)

func InitialConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}
