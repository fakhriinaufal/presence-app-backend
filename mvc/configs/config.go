package configs

import (
	"fmt"
	"os"
	models2 "presence-app-backend/mvc/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	godotenv.Load()

	config := map[string]string{
		"DB_Username": os.Getenv("DB_USERNAME"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
		"DB_Host":     os.Getenv("DB_HOST"),
		"DB_Port":     os.Getenv("DB_PORT"),
		"DB_Name":     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models2.Department{})
	DB.AutoMigrate(&models2.User{})
	DB.AutoMigrate(&models2.Schedule{})
	DB.AutoMigrate(&models2.Presence{})

}
