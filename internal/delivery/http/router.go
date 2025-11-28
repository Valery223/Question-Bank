package http

import (
	v1 "github.com/Valery223/Question-Bank/internal/delivery/http/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *v1.Handler) *gin.Engine {
	router := gin.Default()

	// Можно добавить свои Middleware

	api := router.Group("/api/v1")
	{
		// Questions
		questions := api.Group("/questions")
		questions.POST("/", handler.CreateQuestion)
		questions.GET("/:id", handler.GetQuestionByID)
		// questions.GET("", handler.ListQuestions)
	}

	return router
}
