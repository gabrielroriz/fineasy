package database

import (
	"fmt"
	"os"

	"github.com/gabrielroriz/cli-fineasy/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {

	var err error

	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=fineasy password=%s sslmode=disable",
			os.Getenv("FIN_DB_HOST"),
			os.Getenv("FIN_DB_PORT"),
			os.Getenv("FIN_DB_USER"),
			os.Getenv("FIN_DB_PASSWORD")))

	if err != nil {
		panic("failed to connect database")
	}
	// db.LogMode(true)
	// defer db.Close()

	db.AutoMigrate(
		&models.Source{},
		&models.Wallet{},
		&models.Category{},
		&models.Flow{})
}
