package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"lesson4/lesson5/global"
	"lesson4/lesson5/model"
	"net/http"
	"time"
)

func CreateToken(name string) (string, error) {
	claims := model.MyClaims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "guoyiyang",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(global.SecretSignKey)
	if err != nil {
		fmt.Println("token生成失败")
		return "", err
	}
	return tokenString, nil

}

func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return global.SecretSignKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
func AuthMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "请求头为空",
			})
			c.Abort()
			return
		}
		mc, err := ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "无效的token",
			})
			c.Abort()
			return
		}
		c.Set("name", mc.Name)
		c.Next()

	}
}
