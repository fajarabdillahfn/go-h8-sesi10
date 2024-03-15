package database

import (
	"fmt"
	"log"
	"sesi10/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connection to database:", err)
	}

	fmt.Println("Successfully connected to the database")

	db.AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
