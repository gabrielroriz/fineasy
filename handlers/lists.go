package handlers

import (
	"fineasy/database"
	"fineasy/utils"
)

func ListWallets() {
	values := database.GetWallets()

	if len(values) == 0 {
		return
	}

	var wallets [][]string

	for i := 0; i < len(values); i++ {
		model := (values)[i]
		wallets = append(wallets, model.ToTableFormat())
	}

	utils.TerminalUIPrintTable([]string{"id", "title"}, wallets)

}

func ListSources() {

	values := database.GetSources()

	if len(values) == 0 {
		return
	}

	var sources [][]string

	for i := 0; i < len(values); i++ {
		model := (values)[i]
		sources = append(sources, model.ToTableFormat())
	}

	utils.TerminalUIPrintTable([]string{"id", "title", "flux"}, sources)

}

func ListFlows() {

	values := database.GetFlows()

	if len(*values) == 0 {
		return
	}

	var flows [][]string

	for i := 0; i < len(*values); i++ {
		model := (*values)[i]
		flow := model.ToTableFormat()
		flows = append(flows, flow)
	}

	utils.TerminalUIPrintTable([]string{"id", "date", "source", "flux", "description", "category", "wallet", "cash"}, flows)

}

func ListCategories() {

	values := database.GetCategories()

	if len(values) == 0 {
		return
	}

	var categories [][]string

	for i := 0; i < len(values); i++ {
		model := (values)[i]
		categories = append(categories, model.ToTableFormat())
	}

	utils.TerminalUIPrintTable([]string{"id", "title"}, categories)

}
