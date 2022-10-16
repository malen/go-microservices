package models

type User struct {
	Id       int64  `json:"id" xorm:"pk autoincr"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
