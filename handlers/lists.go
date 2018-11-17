package handlers

import (
	"fmt"

	"github.com/gabrielroriz/fineasy/database"
)

// ListWallets : Shoud have comments here.
func ListWallets() {

	values := database.GetWallets()

	var wallets [][]string

	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		wallet := []string{fmt.Sprintf("%d", model.ID), model.Title}

		wallets = append(wallets, wallet)
	}

	PrintTable([]string{"id", "title"}, wallets)

}

// ListSources : Shoud have comments here.
func ListSources() {

	values := database.GetSources()

	var sources [][]string

	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		source := []string{fmt.Sprintf("%d", model.ID), model.Title, model.Flux}

		sources = append(sources, source)
	}

	PrintTable([]string{"id", "title", "flux"}, sources)

}

// ListFlows : Shoud have description here.
func ListFlows() {

	values := database.GetFlows()

	var total float32

	fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n")
	fmt.Printf(" %-10s | %-20s | %-20s | %-20s | %-20s | %-20s |\n", "id", "date", "source", "category", "wallet", "cash")
	fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n")
	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		total += model.Cash

		date := fmt.Sprintf("%d/%d/%d", model.CreatedAt.Day(), model.CreatedAt.Month(), model.CreatedAt.Year())
		source := fmt.Sprintf("(%d) %s", model.Source.ID, model.Source.Title)
		category := fmt.Sprintf("(%d) %s", model.Category.ID, model.Category.Title)
		wallet := fmt.Sprintf("(%d) %s", model.Wallet.ID, model.Wallet.Title)
		cash := fmt.Sprintf("R$ %.2f", model.Cash)

		fmt.Printf(" %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
			10, fmt.Sprint(model.ID),
			20, date,
			20, source,
			20, category,
			20, wallet,
			20, cash)
	}
	fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n")
	fmt.Printf("                                                                                                        | %-*s |\n", 20, fmt.Sprintf("R$ %.2f", total))
	fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n\n")
}

// ListCategories : Shoud have comments here.
func ListCategories() {

	values := database.GetCategories()

	var categories [][]string

	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		category := []string{fmt.Sprintf("%d", model.ID), model.Title}

		categories = append(categories, category)
	}

	PrintTable([]string{"id", "title"}, categories)

}
