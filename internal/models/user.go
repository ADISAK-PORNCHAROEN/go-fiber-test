package models

type User struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"1234"`
}
