package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}

func (u *User) TableName() string {
	return "users"
}

type UserCreateResp struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type AuthUser struct {
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type ChangePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"'`
}
