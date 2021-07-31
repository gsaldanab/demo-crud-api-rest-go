package commons

import (
	"log"

	"github.com/gsaldanab/demo-crud-api-rest-go/models"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	// dsn := "root:mysql@tcp(go-mysql8:3306)/goSchema?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:mysql@tcp(localhost:3306)/goSchema?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open("mysql", "root:mysql@go-mysql8/goSchema?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer func() {
		dbConn, _ := db.DB()
		dbConn.Close()
	}()

	log.Println("Migrando tablas ...")
	db.AutoMigrate(&models.Persona{})
}
