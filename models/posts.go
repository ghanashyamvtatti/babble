package models

import (
	"time"
)

type Post struct {
	Post      string
	CreatedAt time.Time
	UpdateAt  time.Time
}
