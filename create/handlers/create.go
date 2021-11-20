package handlers

import (
	"encoding/json"

	"github.com/DelusionalOptimist/pastebean/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// CreatePaste creates a new paste and writes back its ID in response
func (h *Handler) CreatePaste(c *gin.Context) {
	p := &models.Paste{}
	json.NewDecoder(c.Request.Body).Decode(p)

	h.HandlerEnv.DB.Create(p)
	c.String(200, "%s", p.ID.String())
}
