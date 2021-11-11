package models

type Paste struct {
	Base
	Body string `json:"body" gorm:"body"`
	Name string `json:"name" gorm:"name"`
}
