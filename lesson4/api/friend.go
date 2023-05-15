package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lesson4/lesson4/model"
	"net/http"
)

func Init() {
	dsn := "root:369248@tcp(127.0.0.1:3306)/csawork?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	model.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	model.DB.AutoMigrate(&model.Friend{})
}
func add(c *gin.Context) {
	var friend model.Friend
	err := c.ShouldBind(&friend)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "shouldbind falied",
		})
	}

	model.DB.Create(&friend)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "add successful",
	})

}
func deleted(c *gin.Context) {
	var friend model.Friend
	name := c.PostForm("name")

	model.DB.Delete(&friend, "name=?", name)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "deleted successful",
	})
}
func findall(c *gin.Context) {
	var friend []model.Friend
	model.DB.Find(&friend)
	c.JSON(http.StatusOK, friend)
}
func grouping(c *gin.Context) {
	var friend model.Friend
	group := c.PostForm("group")
	name := c.PostForm("name")
	model.DB.Model(&friend).Select("group").Where("name=?", name).Update("group", group)
	c.JSON(200, friend)

}
func findone(c *gin.Context) {
	var friend model.Friend
	name := c.PostForm("name")
	model.DB.First(&friend, "name=?", name)
	c.JSON(http.StatusOK, friend)
}
