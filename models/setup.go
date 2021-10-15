package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/joho/godotenv"
	"log"
)

func SetupDB() *gorm.DB {

	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error reading .env file")
	}

	USER := appConfig["USER"]
	PASS := appConfig["PASS"]
	HOST := appConfig["HOST"]
	PORT := appConfig["PORT"]
	DBNAME := appConfig["DBNAME"]
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)


	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}

	return db
}
