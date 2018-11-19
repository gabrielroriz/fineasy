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

	var flows [][]string

	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		id := fmt.Sprintf("%d", model.ID)
		date := fmt.Sprintf("%d/%d/%d", model.CreatedAt.Day(), model.CreatedAt.Month(), model.CreatedAt.Year())
		source := fmt.Sprintf("(%d) %s", model.Source.ID, model.Source.Title)
		flux := fmt.Sprintf("%s", model.Source.Flux)
		description := fmt.Sprintf("%s", model.Description)
		category := fmt.Sprintf("(%d) %s", model.Category.ID, model.Category.Title)
		wallet := fmt.Sprintf("(%d) %s", model.Wallet.ID, model.Wallet.Title)
		cash := fmt.Sprintf("R$ %.2f", model.Cash)

		flow := []string{id, date, source, flux, description, category, wallet, cash}
		flows = append(flows, flow)
	}

	PrintTable([]string{"id", "date", "source", "flux", "description", "category", "wallet", "cash"}, flows)

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
