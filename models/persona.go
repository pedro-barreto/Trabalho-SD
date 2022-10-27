package models

import "github.com/jinzhu/gorm"

type Persona struct {
	gorm.Model

	ID   int64  `json:"id"`
	Nome string `json:"nome"`
}
