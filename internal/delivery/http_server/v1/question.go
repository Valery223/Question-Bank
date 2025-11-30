package v1

import (
	"net/http"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/gin-gonic/gin"
)

// CreateQuestion godoc
// @Summary      Создать новый вопрос
// @Description  Создает вопрос с вариантами ответов. Валидирует входные данные.
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        input body      CreateQuestionRequest  true  "Данные вопроса"
// @Success      201   {object}  map[string]string          "{"status": "created"}"
// @Failure      400   {object}  map[string]string          "Ошибка валидации JSON"
// @Failure      500   {object}  map[string]string          "Внутренняя ошибка сервера"
// @Router       /questions [post]
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

func (h *Handler) UpdateQuestion(c *gin.Context) {
	var req UpdateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	q := req.ToDomain()
	q.ID = domain.ID(c.Param("id"))

	err := h.questionUC.UpdateQuestion(c.Request.Context(), q)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := QuestionToResponse(q)
	c.JSON(http.StatusOK, gin.H{"status": "updated", "question": resp})
}

func (h *Handler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	err := h.questionUC.DeleteQuestion(c.Request.Context(), domain.ID(id))
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
