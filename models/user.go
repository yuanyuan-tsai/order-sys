package models

// login account Email-Password
type User struct {
	ID       string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Orders   []string
}
