package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func InitDB() error {
	dbConnectionString := getDbConnectionString()
	db, err = sql.Open("mysql", dbConnectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to the database")
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed")
	}
}

func getDbConnectionString() string {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_DATABASE")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
}
