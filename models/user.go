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

func GetUserByID(id int64) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, name, email, address, birthdate, city FROM %s WHERE id = ? LIMIT 1", tableName)
	err := db.GetDB().QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.Birthdate, &user.City)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = ? LIMIT 1", tableName)
	err := db.GetDB().QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Address, &user.Birthdate, &user.City)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *User) error {
	insertQuery := fmt.Sprintf("INSERT INTO %s (name, email, password, address, birthdate, city) VALUES (?, ?, ?, ?, ?, ?)", tableName)
	result, err := db.GetDB().Exec(insertQuery, user.Name, user.Email, user.Password, user.Address, user.Birthdate, user.City)
	if err != nil {
		fmt.Println("Insert error:", err)
		return err
	}

	id, _ := result.LastInsertId()
	user.ID = id

	return nil
}

func UpdateUser(user *User) error {
	updateQuery := fmt.Sprintf("UPDATE %s SET name = ?, email = ?, address = ?, birthdate = ?, city = ? WHERE id = ?", tableName)
	_, err := db.GetDB().Exec(updateQuery, user.Name, user.Email, user.Address, user.Birthdate, user.City, user.ID)
	if err != nil {
		fmt.Println("Update error:", err)
		return err
	}
	return nil
}
