package common

import (
	"log"

	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {

	db, err := gorm.Open("mysql", "root:@/teste?charset=utf8")

	if err != nil {

		log.Fatal(err)

	}

	return db

}
