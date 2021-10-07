package database

import (
	"assigment-2/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	host     = "localhost"
	user     = "root"
	password = "Triyono7"
	dbPort   = "3306"
	dbName   = "orders_by"
	db       *gorm.DB
	err      error
)
func StartDB() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, dbPort, dbName)
	dsn := config
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting", err)
	}
	fmt.Println("sukses koneksi ke db")
	err = db.Debug().AutoMigrate(models.Items{}, models.Orders{})
	if err != nil {
		return
	}
}
func GetDB() *gorm.DB {
	return db
}
