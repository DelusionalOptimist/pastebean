package router

import (
	"github.com/DelusionalOptimist/pastebean/update/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := handlers.NewHandler(db)

	r.POST("/update/:id", h.UpdatePaste)

	return r
}
