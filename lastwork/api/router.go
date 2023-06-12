package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lastwork/api/middleware"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.POST("/register", register)
	r.POST("/login", login)
	UserRouter := r.Group("/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.POST("/ask", askQuestion)
		UserRouter.POST("/answer", answerQuestion)
		UserRouter.GET("/queryall", queryAll)
		UserRouter.POST("/get", getQuestion)
	}
	ProblemRouter := UserRouter.Group("/problem")
	{
		ProblemRouter.PUT("/update", updateProblem)
		ProblemRouter.DELETE("/delete", deleteProblem)
	}
	AnswerRouter := UserRouter.Group("/answer")
	{
		AnswerRouter.PUT("/update", updateAnswer)
		AnswerRouter.DELETE("/delete", deleteAnswer)
	}
	r.Run()
}
