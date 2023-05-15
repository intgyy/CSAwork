package model

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type Friend struct {
	gorm.Model
	Name  string `json:"name"`
	Group string `json:"group"`
}
