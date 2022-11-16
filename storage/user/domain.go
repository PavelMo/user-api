package user

import (
	"time"
)

type (
	User struct {
		CreatedAt   time.Time
		DisplayName string
		Email       string
	}
	UserList  map[string]User
	UserStore struct {
		Increment int
		List      UserList
	}
)
