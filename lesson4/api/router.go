package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	Init()
	r := gin.Default()
	r.POST("/add", add)
	r.POST("/deleted", deleted)
	r.GET("/findall", findall)
	r.POST("/grouping", grouping)
	r.POST("/findone", findone)
	r.Run(":8080")
}
