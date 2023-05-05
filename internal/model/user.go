package model

import "time"

type Query struct {
	Username string
}

type User struct {
	Username string
	Email    string
	Password string
	Role     string
}

type UserInfo struct {
	Username  string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
