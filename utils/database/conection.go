package database

import (
	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const DB_USERNAME = "snirdb"
const DB_PASSWORD = "snirdb@2021"
const DB_NAME = "nzimbupaygateway"
const DB_HOST = "172.16.16.36"
const DB_PORT = "5432"

func ConnectionDB() *gorm.DB {

	dsn := "host=" + DB_HOST + " user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " port=" + DB_PORT + " sslmode=disable"

	db, err := gorm.Open("postgres", dsn)

	if err != nil {
		defer db.Close()
		panic("failed to connect database")

	}

	db.AutoMigrate(&domain.Microservice{})

	return db

}
