package api

import (
	"lastwork/api/middleware"
	"lastwork/global"
	"lastwork/model"
	"lastwork/utils"
	"net/http"
	"strconv"

	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"time"
)

func register(c *gin.Context) {
	var user model.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	flag := 0
	global.DB.Model(&model.User{}).Where("username=?", user.Username).Count(&flag)
	if flag != 0 {
		utils.RespFail(c, "user already exists")
		return
	}
	user.Username = username
	user.Password = password
	global.DB.Create(&user)
	utils.RespSuccess(c, "user register successful")

}
func login(c *gin.Context) {
	var user model.User
	var user1 model.User
	err := c.ShouldBindJSON(&user)
	fmt.Println("username" + user.Username)
	fmt.Println("password" + user.Password)
	if err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	var flag *int = new(int)
	global.DB.Model(&model.User{}).Where("username=?", user.Username).Count(flag)
	fmt.Println(*flag)
	if *flag == 0 {
		utils.RespFail(c, "user not exists")
		return
	}

	global.DB.Model(&model.User{}).Where("username=?", user.Username).First(&user1)
	if user.Password != user1.Password {
		utils.RespFail(c, "wrong password")
		return
	}
	claims := model.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "YXH",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(middleware.Secret)

	utils.RespSuccess(c, tokenString)

}
func askQuestion(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "user not exists")
		return
	}
	var user model.User
	global.DB.Model(&model.User{}).Where("username=?", username).First(&user)
	var problem model.Problem
	err := c.ShouldBindJSON(&problem)
	if err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	problem.UserId = user.ID
	global.DB.Model(&model.Problem{}).Create(&problem)
	utils.RespSuccess(c, "user ask problem successful")

}
func answerQuestion(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "user not exists")
		return
	}
	var user model.User
	global.DB.Model(&model.User{}).Where("username=?", username).First(&user)
	var answer model.Answer
	err := c.ShouldBindJSON(&answer)
	if err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	answer.UserId = user.ID
	global.DB.Model(&model.Answer{}).Create(&answer)
	utils.RespSuccess(c, "user answer problem successful")
}
func queryAll(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "user not exists")
		return
	}
	var user model.User
	var problem []model.Problem
	var answer []model.Answer
	global.DB.Model(&model.User{}).Where("username=?", username).First(&user)
	global.DB.Model(&model.Problem{}).Where("user_id=?", user.ID).Find(&problem)
	global.DB.Model(&model.Answer{}).Where("user_id=?", user.ID).Find(&answer)
	c.JSON(http.StatusOK, gin.H{
		"problem": problem,
		"answer":  answer,
	})
}
func updateProblem(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "user not exists")
		return
	}
	var user model.User
	var problem model.Problem
	iid, err := strconv.Atoi(id)
	global.DB.Model(&model.User{}).Where("username=?", username).First(&user)
	fmt.Println(iid)
	err = global.DB.Model(&model.Problem{}).Where("user_id=? AND id=?", user.ID, uint(iid)).First(&problem).Error
	if err != nil {
		utils.RespFail(c, "problem not exists")
		fmt.Println(err)
		return
	}
	c.ShouldBindJSON(&problem)
	global.DB.Model(&model.Problem{}).Save(&problem)
	utils.RespSuccess(c, "update problem successful")
}
func updateAnswer(c *gin.Context) {
	id := c.Query("id")
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "user not exists")
		return
	}
	var user model.User
	var answer model.Answer

	global.DB.Model(&model.User{}).Where("username=?", username).First(&user)
	iid, err := strconv.Atoi(id)
	err = global.DB.Model(&model.Answer{}).Where("user_id=? AND id=?", user.ID, uint(iid)).First(&answer).Error
	if err != nil {

		utils.RespFail(c, "problem not exists")
		fmt.Println(err.Error())
		return
	}
	c.ShouldBindJSON(&answer)
	global.DB.Model(&model.Answer{}).Save(&answer)
	utils.RespSuccess(c, "update answer successful")
}
func deleteProblem(c *gin.Context) {
	id := c.Query("id")
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "user not exists")
		return
	}
	var user model.User
	var problem model.Problem
	iid, err := strconv.Atoi(id)
	global.DB.Model(&model.User{}).Where("username=?", username).First(&user)
	err = global.DB.Model(&model.Problem{}).Where("user_id=? AND id=?", user.ID, uint(iid)).First(&problem).Error
	if err != nil {
		utils.RespFail(c, "problem not exists")
		return
	}
	global.DB.Model(&model.Problem{}).Delete(&problem)
	utils.RespSuccess(c, "delete problem successful")
}
func deleteAnswer(c *gin.Context) {
	id := c.Query("id")
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "user not exists")
		return
	}
	var user model.User
	var answer model.Answer
	iid, err := strconv.Atoi(id)
	global.DB.Model(&model.User{}).Where("username=?", username).First(&user)
	err = global.DB.Model(&model.Answer{}).Where("user_id=? AND id=?", user.ID, uint(iid)).First(&answer).Error
	if err != nil {
		utils.RespFail(c, "answer not exists")
		return
	}
	global.DB.Model(&model.Answer{}).Delete(&answer)
	utils.RespSuccess(c, "delete answer successful")
}
func getQuestion(c *gin.Context) {
	id := c.Query("id")
	iid, err := strconv.Atoi(id)
	var problem model.Problem
	var answer model.Answer
	err = global.DB.Model(&model.Problem{}).Where("id=?", uint(iid)).First(&problem).Error
	if err != nil {
		utils.RespFail(c, "problem not exists")
		return
	}
	err = global.DB.Model(&model.Answer{}).Where("problem_id=?", uint(iid)).First(&answer).Error
	c.JSON(http.StatusOK, gin.H{
		"problem": problem,
		"answer":  answer,
	})
}
