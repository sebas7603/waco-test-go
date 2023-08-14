package config

import (
	"github.com/joho/godotenv"
)

func InitialConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	// TODO: Create tables in DB

	return nil
}
