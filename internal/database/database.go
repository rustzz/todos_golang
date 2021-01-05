package database

import (
	"log"
	"os"

	"github.com/rustzz/todos/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDatabase : ...
func ConnectDatabase() *gorm.DB {
	var dsn = os.Getenv("DB_DSN")
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return db
}

// Migrate : ...
func Migrate() {
	var db = ConnectDatabase()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Note{})
	var dbConn, _ = db.DB()
	dbConn.Close()
	return
}
