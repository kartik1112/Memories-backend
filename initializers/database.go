package initializers

import (
	"fmt"
	"os"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_NAME")
	// sslMode := os.Getenv("MYSQL_SSLMODE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	DB.Exec("CREATE DATABASE memories;")
	DB.Exec("USE memories;")

	if err != nil {
		log.Fatal("Failed to connect to Database..")
	}

}
