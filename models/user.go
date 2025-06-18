package models

type User struct {
	ID       string
	Name     string
	Email    string // login account
	Password string
	Role     string
	Orders   []string
}
