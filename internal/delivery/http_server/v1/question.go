package v1

import (
	"net/http"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateQuestion(c *gin.Context) {
	//  Биндинг + Валидация формата (Gin)
	var req CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	//  Вызов UseCase
	// Превращаем DTO в Domain и передаем контекст
	var q = req.ToDomain()
	err := h.questionUC.CreateQuestion(c.Request.Context(), q)

	if err != nil {
		// Тут можно проверить ошибку через errors.Is(err, domain.ErrInvalidDifficulty)
		// и вернуть 400 вместо 500. Но для простоты:
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var resp = QuestionToResponse(q)
	c.JSON(http.StatusCreated, gin.H{"status": "created", "question": resp})
}

func (h *Handler) GetQuestionByID(c *gin.Context) {
	// Получаем ID из пути
	id := c.Param("id")

	// Вызов UseCase
	q, err := h.questionUC.GetQuestionByID(c.Request.Context(), domain.ID(id))
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := QuestionToResponse(q)
	c.JSON(http.StatusOK, resp)
}
