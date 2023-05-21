package api

import (
	"github.com/gin-gonic/gin"
	"lesson4/lesson5/global"
	"lesson4/lesson5/jwt"
	"lesson4/lesson5/model"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	count := 0
	global.DB.Model(&model.User{}).Where("name=?", user.Name).Count(&count)
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户已存在",
		})
		return
	}
	global.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})

}
func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	count := 0
	global.DB.Model(&model.User{}).Where("name=?", user.Name).Count(&count)
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户未存在",
		})
		return
	}
	tokenString, err := jwt.CreateToken(user.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "token生成失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": tokenString,
	})

}
func GetUser(c *gin.Context) {
	var user model.User
	var err error
	userName := c.MustGet("name").(string)
	user.Password, err = global.Rbq.Get(c, userName).Result()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"name":     userName,
			"password": user.Password,
		})
		return
	}

	global.DB.Where("name=?", userName).First(&user)

	global.Rbq.Set(c, userName, user.Password, time.Minute)

	c.JSON(200, gin.H{
		"msg": user,
	})

}
