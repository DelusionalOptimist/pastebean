package router

import (
	"github.com/DelusionalOptimist/pastebean/delete/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	h := handlers.NewHandler(db)

	r.DELETE("/delete/:id", h.DeletePaste)

	return r
}
