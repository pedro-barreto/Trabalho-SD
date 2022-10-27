package commons

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pedro-barreto/Trabalho-SD/models"
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

	log.Println("Migrando....")

	db.AutoMigrate(&models.Persona{})
}
