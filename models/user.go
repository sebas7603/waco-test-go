package models

import (
	"fmt"
	"time"

	"github.com/sebas7603/waco-test-go/pkg/db"
)

var tableName = "users"

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Address   string
	Birthdate time.Time
	City      string
}

func GetAllUsers() (*[]User, error) {
	var users []User
	query := fmt.Sprintf("SELECT id, name, email, address, birthdate, city FROM %s", tableName)
	rows, err := db.GetDB().Query(query)
	if err != nil {
		fmt.Println("Query error:", err)
		return nil, err
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.Birthdate, &user.City); err != nil {
			fmt.Println("Scan error:", err)
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}
