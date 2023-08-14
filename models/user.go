package models

import (
	"time"
)

type User struct {
	ID        int64
	name      string
	email     string
	password  string
	address   string
	birthdate time.Time
	city      string
}
