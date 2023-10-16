package model

type User struct {
	Id       int64  `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
