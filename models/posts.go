package models

import (
	"time"
)

type Post struct {
	Id        int
	UserId    int
	Post      string
	CreatedAt time.Time
	UpdateAt  time.Time
}
