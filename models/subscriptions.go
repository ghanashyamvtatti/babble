package models

import (
	"time"
)

type Subscription struct {
	Id         int
	Subscriber *User
	Publisher  *User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
