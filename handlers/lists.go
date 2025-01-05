package handlers

import (
	"fineasy/database"
	"fineasy/utils"
)

func ListWallets(selectionMode bool) *string {
	values := database.GetWallets()

	if len(values) == 0 {
		return nil
	}

	var wallets [][]string

	for i := 0; i < len(values); i++ {
		model := (values)[i]
		wallets = append(wallets, model.ToTableFormat())
	}

	tableKeys := []string{"id", "title"}
	tableValues := wallets

	if selectionMode {
		return utils.TerminalUIPrintTableSelectionMode(tableKeys, tableValues)
	} else {
		utils.TerminalUIPrintTableModeOnlyView(tableKeys, tableValues)
		return nil
	}

}

func ListSources(selectionMode bool) *string {

	values := database.GetSources()

	if len(values) == 0 {
		return nil
	}

	var sources [][]string

	for i := 0; i < len(values); i++ {
		model := (values)[i]
		sources = append(sources, model.ToTableFormat())
	}

	tableKeys := []string{"id", "title", "flux"}
	tableValues := sources

	if selectionMode {
		return utils.TerminalUIPrintTableSelectionMode(tableKeys, tableValues)
	} else {
		utils.TerminalUIPrintTableModeOnlyView(tableKeys, tableValues)
		return nil
	}

}

func ListFlows(selectionMode bool) *string {

	values := database.GetFlows()

	if len(*values) == 0 {
		return nil
	}

	var flows [][]string

	for i := 0; i < len(*values); i++ {
		model := (*values)[i]
		flow := model.ToTableFormat()
		flows = append(flows, flow)
	}

	tableKeys := []string{"id", "date", "source", "flux", "description", "category", "wallet", "cash"}
	tableValues := flows

	if selectionMode {
		return utils.TerminalUIPrintTableSelectionMode(tableKeys, tableValues)
	} else {
		utils.TerminalUIPrintTableModeOnlyView(tableKeys, tableValues)
		return nil
	}

}

func ListCategories(selectionMode bool) *string {

	values := database.GetCategories()

	if len(values) == 0 {
		return nil
	}

	var categories [][]string

	for i := 0; i < len(values); i++ {
		model := (values)[i]
		categories = append(categories, model.ToTableFormat())
	}

	tableKeys := []string{"id", "title"}
	tableValues := categories

	if selectionMode {
		return utils.TerminalUIPrintTableSelectionMode(tableKeys, tableValues)
	} else {
		utils.TerminalUIPrintTableModeOnlyView(tableKeys, tableValues)
		return nil
	}
}
