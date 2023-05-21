package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`

	Password string `json:"password"`
}
type MyClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
