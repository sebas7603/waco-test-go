package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sebas7603/waco-test-go/pkg/db"
)

var err error

func InitialConfig() error {
	err = godotenv.Load(".env")
	if err != nil {
		return err
	}

	err = db.InitDB()
	if err != nil {
		fmt.Println("Error connecting to DB", err)
		return err
	}

	err = createTablesInDB()
	if err != nil {
		return err
	}

	return nil
}

func createTablesInDB() error {
	err = db.CreateUsersTable()
	if err != nil {
		fmt.Println("Error creating users table", err)
		return err
	}

	err = db.CreateFavoritesTable()
	if err != nil {
		fmt.Println("Error creating favorites table", err)
		return err
	}

	return nil
}
