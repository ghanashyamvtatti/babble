package models

import (
	"time"
)

type User struct {
	FullName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
