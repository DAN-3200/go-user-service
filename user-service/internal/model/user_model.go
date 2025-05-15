package model

type User struct {
	Id       int    `json:"id"` // temporáriamente
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Date     string `json:"date"`
}
