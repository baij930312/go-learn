package model

type User struct {
	UserId   int `json:"userId"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}
