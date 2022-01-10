package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"id"`
	CreatedAt time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
}
