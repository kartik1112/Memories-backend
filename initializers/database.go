package initializers

import (
	"os"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	dsn := os.Getenv("DB_CONNECTION_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Database..")
	}

}
