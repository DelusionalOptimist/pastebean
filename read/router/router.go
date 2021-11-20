package router

import (
	"github.com/DelusionalOptimist/pastebean/read/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := handlers.NewHandler(db)

	r.GET("/:id", h.GetPaste)

	return r
}
