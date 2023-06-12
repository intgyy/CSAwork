package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"username" json:"password" binding:"required"`
	gorm.Model
}
type Problem struct {
	gorm.Model
	UserId uint   `form:"userId" json:"userId"`
	Title  string `form:"title" json:"title" binding:"required"`

	Content string `form:"content" json:"content" binding:"required"`
}
type Answer struct {
	gorm.Model
	UserId    uint   `form:"userId" json:"userId"`
	ProblemId uint   `form:"problemId" json:"problemId" binding:"required"`
	Content   string `form:"content" json:"content" binding:"required"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
