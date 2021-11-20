package models

import (
	"gorm.io/gorm"
)

type HandlerEnv struct {
	DB *gorm.DB
}
