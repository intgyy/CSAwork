package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lesson3/api/middleware"
	"lesson3/dao"
	"lesson3/model"
	"lesson3/utils"
	"net/http"
	"time"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if flag {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "user already exists",
		})
		return
	}
	dao.Adduser(username, password)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "add user successful",
	})

}
func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := dao.SelectUser(username)
	if !flag {
		utils.RespFail(c, "user doesn`t exists")
		return
	}
	seletpassword := dao.SelectPasswordFromUsername(username)
	if seletpassword != password {
		utils.RespFail(c, "wrong password")
		return
	}
	claims := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "YXH",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}
