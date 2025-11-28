package httpServer

import (
	v1 "github.com/Valery223/Question-Bank/internal/delivery/http_server/v1"
	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *v1.Handler) *gin.Engine {
	router := gin.Default()

	// Можно добавить свои Middleware
	router.Use(AuthMiddleware())
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Здесь проверяем JWT, сессию и т.д.
		userRole := c.GetHeader("User-Role")
		userID := c.GetHeader("User-ID")

		//  Сохраняем в gin.Context
		c.Set("userID", userID)
		c.Set("userRole", userRole)
		// ИЛИ сохраняем в стандартный context
		ctx := domain.NewContextWithUser(c.Request.Context(), domain.ID(userID), domain.UserRole(userRole))
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
