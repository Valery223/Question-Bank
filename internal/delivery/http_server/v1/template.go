package v1

import (
	"net/http"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/gin-gonic/gin"
)

// Создание шаблона теста
//
// FIXME: Решить, что именно возвращать
// Принимает CreateTemplateRequest, возвращает TemplatesResponse (с ID вопросов)
// Либо можно возвращать полные вопросы TemplateDetailsResponse, но это может быть избыточно
func (h *Handler) CreateTemplate(c *gin.Context) {
	//  Биндинг + Валидация формата (Gin)
	var req CreateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid request body, "+err.Error())
		return
	}

	//  Вызов UseCase
	// Превращаем DTO в Domain и передаем контекст
	var t = req.ToDomain()
	err := h.templateUC.CreateTemplate(c.Request.Context(), t)

	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "failed to create template, "+err.Error())
		return
	}
	var resp = TemplateToResponse(t)
	c.JSON(http.StatusCreated, gin.H{"status": "created", "template": resp})

}

func (h *Handler) GetTemplateByID(c *gin.Context) {
	id := c.Param("id")
	template, err := h.templateUC.GetTemplateByID(c.Request.Context(), domain.ID(id))
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "failed to get template, "+err.Error())
		return
	}
	var resp = TemplateToResponse(template)
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateTemplate(c *gin.Context) {
	var req UpdateTemplateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.errorResponse(c, http.StatusBadRequest, "invalid request body, "+err.Error())
		return
	}

	t := req.ToDomain()
	t.ID = domain.ID(c.Param("id"))

	err := h.templateUC.UpdateTemplate(c.Request.Context(), t)
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "failed to update template, "+err.Error())
		return
	}

	resp := TemplateToResponse(t)
	c.JSON(http.StatusOK, gin.H{"status": "updated", "template": resp})
}

func (h *Handler) DeleteTemplate(c *gin.Context) {
	id := c.Param("id")
	err := h.templateUC.DeleteTemplate(c.Request.Context(), domain.ID(id))
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, "failed to delete template, "+err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
