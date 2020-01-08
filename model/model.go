package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Level struct {
	ID   int    `gorm:"column:id"`
	Rate int    `gorm:"column:rate"`
	Code string `gorm:"column:code"`
	Name string `gorm:"column:name"`
	Tag  string `gorm:"column:tag"`
}

func (l Level) TableName() string {
	return "truck_level"
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type UserClaims struct {
	UserID int    `json:"uid"`
	Phone  string `json:"p"`
	Email  string `json:"e"`
	jwt.StandardClaims
}

type Error struct {
	Code int
	Msg  string
}

type UserUpdateReq struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResp struct {
	*User
	Token string `json:"token"`
}
