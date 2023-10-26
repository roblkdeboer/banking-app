package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Password  string
}

type UserSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}