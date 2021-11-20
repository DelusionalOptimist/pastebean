package handlers

import (
	"encoding/json"

	"github.com/DelusionalOptimist/pastebean/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// GetPaste retrieves a paste with the ID given as paramter
func (h *Handler) GetPaste(c *gin.Context) {
	id := c.Param("id")
	p := &models.Paste{}

	h.HandlerEnv.DB.First(p, "id = ?", id)
	json.NewEncoder(c.Writer).Encode(p)
}
