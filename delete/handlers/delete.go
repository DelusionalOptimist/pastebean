package handlers

import (
	"github.com/DelusionalOptimist/pastebean/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// DeletePaste deletes a paste with the ID given as paramter
func (h *Handler) DeletePaste(c *gin.Context) {
	id := c.Param("id")
	p := &models.Paste{}

	h.HandlerEnv.DB.Delete(p, "ID = ?", id)
}
