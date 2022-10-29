package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pedro-barreto/Trabalho-SD/backend/models"
)

func getAllAlunos() {

	aluno := models.Aluno{}

	db, err := gorm.Open("mysql", "root:@/escola?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.DB().Query("SELECT * FROM alunos")

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		err := res.Scan(&aluno.ID, &aluno.Nome, &aluno.Idade, &aluno.Telefone)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", aluno)
	}

}

func main() {

	getAllAlunos()

}
