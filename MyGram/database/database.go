import (
	"fmt"
	"log"
	"os"
	"mygram/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("MYSQL_HOST")
	port     = os.Getenv("MYSQL_PORT")
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	dbname   = os.Getenv("MYSQL_DATABASE")

	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database:", err)
	}

	fmt.Println("Successfully connected to the database")

	db.Debug().AutoMigrate(&models.User{}, &models.Comment{}, &models.Photo{}, &models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
