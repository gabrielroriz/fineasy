package handlers

import (
	"fmt"

	db "github.com/gabrielroriz/cli-fineasy/database"
)

func ListFlows() bool {

	values := db.GetFlows()

	var total float32 = 0

	fmt.Printf("-------------------------------------------------------------------------------------------------------------------\n")
	fmt.Printf(" %-20s | %-20s | %-20s | %-20s | %-20s |\n", "date", "source", "category", "wallet", "cash")
	fmt.Printf("-------------------------------------------------------------------------------------------------------------------\n")
	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		total += model.Cash

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
	fmt.Printf("-------------------------------------------------------------------------------------------------------------------\n")
	fmt.Printf("                                                                                           | %-*s |\n", 20, fmt.Sprintf("R$ %.2f", total))
	fmt.Printf("-------------------------------------------------------------------------------------------------------------------\n")
	return true
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
