package model

type Query struct {
	Username string
}

type User struct {
	Username string
	Email    string
	Password string
	Role     string
}
