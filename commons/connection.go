package commons

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pedro-barreto/Trabalho-SD/models"
)

func GetConnection() *gorm.DB {

	db, err := gorm.Open("mysql", "root:@/teste?charset=utf8")

	if err != nil {

		log.Fatal(err)

	}

	return db

}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando....")

	db.AutoMigrate(&models.Persona{})
}
