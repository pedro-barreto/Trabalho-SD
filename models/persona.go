package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Persona struct {
	gorm.Model

	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nome      string `json:"nome"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
