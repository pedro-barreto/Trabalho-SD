package models

type Aluno struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nome     string `json:"nome"`
	Idade    int64  `json:"idade"`
	Telefone string `json:"telefone"`
}
