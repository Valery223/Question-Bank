package v1

import (
	"net/http"

	"github.com/Valery223/Question-Bank/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSession(c *gin.Context) {
	var req CreateTestSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	session := req.ToDomain()

	err := h.sessionUC.CreateSession(c.Request.Context(), session)
	if err != nil {
		h.errorResponse(c, 500, err.Error())
		return
	}

	res := TestSessionToResponse(session)

	c.JSON(http.StatusCreated, gin.H{"status": "created", "session": res})
}

func (h *Handler) GetSessionByID(c *gin.Context) {
	id := c.Param("id")

	session, err := h.sessionUC.GetSessionByID(c.Request.Context(), domain.ID(id))
	if err != nil {
		h.errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	res := TestSessionToResponse(session)
	c.JSON(http.StatusOK, res)
}
