package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lesson4/lesson5/api"
	"lesson4/lesson5/global"
	"lesson4/lesson5/jwt"
)

func main() {
	global.Initredis()
	global.Initmysql()
	r := gin.Default()

	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	group := r.Group("/user").Use(jwt.AuthMiddleWare())
	group.GET("get", api.GetUser)
	r.Run(":8080")
}
