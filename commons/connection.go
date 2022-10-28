package commons

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pedro-barreto/Trabalho-SD/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/test?charset=utf8")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Executando o banco....")

	db.AutoMigrate(&models.Aluno{})
}
