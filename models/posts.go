package models

import (
	"time"
)

type Post struct {
	Post      string
	Username  string
	CreatedAt time.Time
	UpdateAt  time.Time
}
