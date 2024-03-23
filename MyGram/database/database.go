package database

import (
	"fmt"
	"log"
	"mygram/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_DATABASE")

	db  *gorm.DB
	err error
)

func StartDB() {
	// Database connection configuration for MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// Open a connection to MySQL
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Auto migrate model tables
	err = db.AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
	if err != nil {
		log.Fatal("Failed to auto migrate tables: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
