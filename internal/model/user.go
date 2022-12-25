package model

type User struct {
	ID       *int
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
