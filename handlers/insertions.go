package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"fineasy/database"
	"fineasy/utils"
)

func InsertWallet() {

	wallet := database.Wallet{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nYou're creating a new wallet:")
	utils.TerminalUIPrintBold("What is the title? ")
	text, _ := reader.ReadString('\n')

	wallet.Title = strings.Split(text, "\n")[0]

	if err := database.InsertWallet(&wallet); err != nil {
		fmt.Println(err)
	} else {
		utils.TerminalUIPrintGreen("New wallet added successfuly.\n")
	}

}

func InsertSource() {

	source := database.Source{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nYou're creating a new source:")

	utils.TerminalUIPrintBold("What is the title? ")
	input, _ := reader.ReadString('\n')
	source.Title = strings.Split(input, "\n")[0]

	utils.TerminalUIPrintBold("Enter [1] 'expense' or [2] 'income': ")
	fmt.Scanf("%s", &input)

	if input == "1" {
		source.Flux = "expense"
	} else if input == "2" {
		source.Flux = "income"
	}

	if err := database.InsertSource(&source); err != nil {
		fmt.Println(err)
	} else {
		utils.TerminalUIPrintGreen("New source added successfuly.\n")
	}

}

func InsertCategory() {

	category := database.Category{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nYou're creating a new category:")
	utils.TerminalUIPrintBold("What is the title? ")
	text, _ := reader.ReadString('\n')

	category.Title = strings.Split(text, "\n")[0]

	if err := database.InsertCategory(&category); err != nil {
		fmt.Println(err)
	} else {
		utils.TerminalUIPrintGreen("New category added successfuly.\n")
	}

}

func InsertFlow() {
	utils.TerminalClearScreen()
	ListCategories()
	category := selectModel(database.ConvertToModels(database.GetCategories())).(database.Category)

	utils.TerminalClearScreen()
	ListWallets()
	wallet := selectModel(database.ConvertToModels(database.GetWallets())).(database.Wallet)

	utils.TerminalClearScreen()
	ListSources()
	source := selectModel(database.ConvertToModels(database.GetSources())).(database.Source)

	var answer string

	// Description
	utils.TerminalClearScreen()
	reader := bufio.NewReader(os.Stdin)
	utils.TerminalUIPrintBold("What is the description? ")
	text, _ := reader.ReadString('\n')

	description := strings.Split(text, "\n")[0]

	// Cash
	utils.TerminalClearScreen()
	var cash float32
	utils.TerminalUIPrintBold("How much money? ")
	fmt.Scanf("%f", &cash)

	// Last Confirmation
	utils.TerminalClearScreen()

	flow := database.Flow{
		CategoryID:  category.GetID(),
		WalletID:    wallet.GetID(),
		SourceID:    source.GetID(),
		Description: description,
		Cash:        cash,
	}

	utils.TerminalUIPrintTable([]string{"categoryId", "walletId", "sourceId", "description", "cash"}, [][]string{flow.InMemoryTableFormat()})

	fmt.Printf("\n\nConfirm flow insertion? Y or N: ")
	fmt.Scanf("%s", &answer)

	if answer == "Y" {
		if err := database.InsertFlow(&flow); err != nil {
			fmt.Println(err)
		} else {
			utils.TerminalUIPrintGreen("\nNew flow added successfuly.\n")
		}
	}
}

func selectModel(list []database.Model) database.Model {

	var answer string
	var id uint

	for answer != "Y" {

		fmt.Print(utils.MakeBold("What is %s ID? ", list[0].GetTypeInString()))
		fmt.Scanf("%d", &id)

		var model database.Model

		for i := 0; i < len(list); i++ {

			model = (list)[i]

			if model.GetID() == id {
				utils.TerminalPrintOnSameLine("Confirm that is %s %s? Type Y (enter) or N: ", model.GetTypeInString(), utils.MakeBold(model.ToString()))
				fmt.Scanf("%s", &answer)
				if len(answer) == 0 {
					answer = "Y"
				}
				break
			}
		}

		if answer == "Y" {
			fmt.Printf("%s: %s.\n", model.GetTypeInString(), utils.MakeBold(model.ToString()))
			return model
		}

	}

	return nil

}
