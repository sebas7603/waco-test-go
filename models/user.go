package models

import (
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Address   string
	Birthdate time.Time
	City      string
}
