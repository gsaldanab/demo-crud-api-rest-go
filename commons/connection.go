package commons

import (
	"log"

	"github.com/gsaldanab/demo-crud-api-rest-go/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:mysql@/goSchema?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando tablas ...")
	db.AutoMigrate(&models.Persona{})
}
