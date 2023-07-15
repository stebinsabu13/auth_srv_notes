package models

type User struct {
	Id       int64  `json:"id" gorm:"primarykey;auto_increment"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
