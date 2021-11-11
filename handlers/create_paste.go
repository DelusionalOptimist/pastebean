package handlers

import (
	"encoding/json"

	"github.com/DelusionalOptimist/pastebean/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePaste(c *gin.Context) {
	p := models.Paste{}
	json.NewDecoder(c.Request.Body).Decode(&p)

	h.HandlerEnv.DB.Create(&p)
}
