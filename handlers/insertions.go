package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gabrielroriz/fineasy/database"
)

func InsertWallet() {

	wallet := database.Wallet{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nYou're creating a new wallet:")
	PrintBold("What is the title? ")
	text, _ := reader.ReadString('\n')

	wallet.Title = strings.Split(text, "\n")[0]

	if err := database.InsertWallet(&wallet); err != nil {
		fmt.Println(err)
	} else {
		PrintSuccess("New wallet added successfuly.\n")
	}

}

func InsertSource() {

	source := database.Source{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nYou're creating a new source:")

	PrintBold("What is the title? ")
	input, _ := reader.ReadString('\n')
	source.Title = strings.Split(input, "\n")[0]

	PrintBold("Enter [1] 'expense' or [2] 'income': ")
	fmt.Scanf("%s", &input)

	if input == "1" {
		source.Flux = "expense"
	} else if input == "2" {
		source.Flux = "income"
	}

	if err := database.InsertSource(&source); err != nil {
		fmt.Println(err)
	} else {
		PrintSuccess("New source added successfuly.\n")
	}

}

func InsertCategory() {

	category := database.Category{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nYou're creating a new category:")
	PrintBold("What is the title? ")
	text, _ := reader.ReadString('\n')

	category.Title = strings.Split(text, "\n")[0]

	if err := database.InsertCategory(&category); err != nil {
		fmt.Println(err)
	} else {
		PrintSuccess("New category added successfuly.\n")
	}

}

func InsertFlow() {

	category := selectModel(database.ConvertToModels(database.GetCategories())).(database.Category)
	wallet := selectModel(database.ConvertToModels(database.GetWallets())).(database.Wallet)
	source := selectModel(database.ConvertToModels(database.GetSources())).(database.Source)

	var answer string

	//description
	reader := bufio.NewReader(os.Stdin)

	PrintBold("What is the description? ")
	text, _ := reader.ReadString('\n')

	description := strings.Split(text, "\n")[0]
	fmt.Printf("\033[1A\033[KDescription: \033[1m%s.\033[0m\n", description)

	//cash
	var cash float32

	PrintBold("How much money? ")
	fmt.Scanf("%f", &cash)
	fmt.Printf("\033[1A\033[KCash: \033[1m%.2f.\033[0m", cash)

	fmt.Printf("\n\nConfirm flow insertion? Y or N: ")
	fmt.Scanf("%s", &answer)

	if answer == "Y" {

		if err := database.InsertFlow(
			&database.Flow{
				CategoryID:  category.GetID(),
				WalletID:    wallet.GetID(),
				SourceID:    source.GetID(),
				Description: description,
				Cash:        cash,
			}); err != nil {
			fmt.Println(err)
		} else {
			PrintSuccess("\nNew flow added successfuly.\n")
		}
	}
}

func selectModel(list []database.Model) database.Model {

	var answer string
	var id uint

	for answer != "Y" {

		if id > 0 {
			fmt.Print(MakeBold(fmt.Sprintf("\033[1A\033[KWhat is %s ID? ", list[0].GetTypeInString())))
		} else {
			fmt.Print(MakeBold(fmt.Sprintf("\033[KWhat is %s ID? ", list[0].GetTypeInString())))
		}

		fmt.Scanf("%d", &id)

		var model database.Model

		for i := 0; i < len(list); i++ {

			model = (list)[i]

			if model.GetID() == id {
				fmt.Printf("\033[1A\033[KConfirm that is %s %s? Type Y or N: ", model.GetTypeInString(), MakeBold(model.ToString()))
				fmt.Scanf("%s", &answer)
				break
			}
		}

		if answer == "Y" {
			fmt.Printf("\033[1A\033[K%s: %s.\n", model.GetTypeInString(), MakeBold(model.ToString()))
			return model
		}

	}

	return nil

}
