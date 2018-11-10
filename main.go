package main

import (
	"fmt"
	"os"

	"github.com/gabrielroriz/cli-fineasy/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=fineasy password=%s sslmode=disable",
			os.Getenv("FIN_DB_HOST"),
			os.Getenv("FIN_DB_PORT"),
			os.Getenv("FIN_DB_USER"),
			os.Getenv("FIN_DB_PASSWORD")))

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(
		&models.Source{},
		&models.Wallet{},
		&models.Category{},
		&models.Flow{})

	// Create

	// db.Create(&models.Category{Title: "ENTRETENIMENTO"})
	// db.Create(&models.Wallet{Title: "BB"})
	// db.Create(&models.Source{Title: "Café RA", Flux: "expense"})

	db.Create(&models.Flow{
		Cash:       1500,
		SourceID:   1,
		WalletID:   1,
		CategoryID: 1,
		Flux:       "expense"})

	// db.Create()

	// db.Create(&models.Category{Title: "ENTRETENIMENTO"})
	// db.Create(&models.Category{Title: "CONTAS REPSOL"})

	// // Read
	// var product Product
	// db.First(&product, 1)                   // find product with id 1
	// db.First(&product, "code = ?", "L1212") // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&product)
}
