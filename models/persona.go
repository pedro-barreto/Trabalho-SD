package models

type Persona struct {
	ID   int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nome string `json:"nome"`
}
