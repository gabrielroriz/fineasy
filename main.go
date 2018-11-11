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
	// db.LogMode(true)
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(
		&models.Source{},
		&models.Wallet{},
		&models.Category{},
		&models.Flow{})

	// Create

	// argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	if len(argsWithoutProg) != 0 && argsWithoutProg[0] != "" {
		switch argsWithoutProg[0] {
		case "-lf":
			res := db.Preload("Source").Preload("Category").Preload("Wallet").Find(&[]models.Flow{})

			// fmt.Println(reflect.TypeOf((res.Value)))

			if value, ok := res.Value.(*[]models.Flow); ok {
				fmt.Printf("-------------------------------------------------------------------------------------------------------------------\n")
				fmt.Printf(" %-20s | %-20s | %-20s | %-20s | %-20s |\n", "date", "source", "category", "wallet", "cash")
				fmt.Printf("-------------------------------------------------------------------------------------------------------------------\n")
				for i := 0; i < len(*value); i++ {

					model := db.Model((*value)[i])

					// fmt.Println(reflect.TypeOf(model.Value))

					if model, ok := model.Value.(models.Flow); ok {

						date := fmt.Sprintf("%d/%d/%d", model.CreatedAt.Day(), model.CreatedAt.Month(), model.CreatedAt.Year())
						source := fmt.Sprintf("(%d) %s", model.Source.ID, model.Source.Title)
						category := fmt.Sprintf("(%d) %s", model.Category.ID, model.Category.Title)
						wallet := fmt.Sprintf("(%d) %s", model.Wallet.ID, model.Wallet.Title)
						cash := fmt.Sprintf("R$ %.2f", model.Cash)

						fmt.Printf(" %-*s | %-*s | %-*s | %-*s | %-*s |\n",
							20, date,
							20, source,
							20, category,
							20, wallet,
							20, cash)
					}

				}
				fmt.Printf("-------------------------------------------------------------------------------------------------------------------\n")

			} else {
				fmt.Println("ERROR")
			}
		default:

		}
	}

	// db.Create(&models.Category{Title: "ENTRETENIMENTO"})
	// db.Create(&models.Wallet{Title: "BB"})
	// db.Create(&models.Source{Title: "Café RA", Flux: "expense"})

	// db.Create(&models.Flow{
	// 	Cash:       8000,
	// 	SourceID:   1,
	// 	WalletID:   1,
	// 	CategoryID: 1,
	// 	Flux:       "expense"}).Association("sources")

	// db.Create(&models.Flow{
	// 	Cash:       204,
	// 	SourceID:   1,
	// 	WalletID:   1,
	// 	CategoryID: 1,
	// 	Flux:       "expense"})

	// db.Create(&models.Flow{
	// 	Cash:       180,
	// 	SourceID:   1,
	// 	WalletID:   1,
	// 	CategoryID: 1,
	// 	Flux:       "expense"})

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
