package db

import (
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUsersTable() error {
	usersSQL, err := ioutil.ReadFile("sql/create_users_table.sql")
	if err != nil {
		return err
	}

	err = CreateTable(string(usersSQL))
	if err != nil {
		return err
	}

	return nil
}

func CreateFavoritesTable() error {
	favoritesSQL, err := ioutil.ReadFile("sql/create_favorites_table.sql")
	if err != nil {
		return err
	}

	err = CreateTable(string(favoritesSQL))
	if err != nil {
		return err
	}

	return nil
}
