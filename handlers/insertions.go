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

	categories := database.GetCategories()
	wallets := database.GetWallets()
	sources := database.GetSources()

	var categoryId uint

	var answer string

	for answer != "Y" {
		if categoryId > 0 {
			PrintBold("\033[1A\033[KWhat is category id? ")
		} else {
			PrintBold("\033[KWhat is category id? ")
		}

		fmt.Scanf("%d", &categoryId)

		var model database.Category

		for i := 0; i < len(*categories); i++ {

			model = (*categories)[i]

			if model.ID == categoryId {
				fmt.Printf("\033[1A\033[KConfirm that is \033[1mCategory (%d) %s\033[0m? Type Y or N: ", model.ID, model.Title)
				fmt.Scanf("%s", &answer)
				break
			}
		}

		if answer == "Y" {
			fmt.Printf("\033[1A\033[KCategory: \033[1m(%d) %s.\033[0m", model.ID, model.Title)
		}

	}

	answer = ""
	var walletId uint

	for answer != "Y" {
		if walletId > 0 {
			PrintBold("\033[1A\033[KWhat is wallet id? ")
		} else {
			PrintBold("\nWhat is wallet id? ")
		}

		fmt.Scanf("%d", &walletId)

		var model database.Wallet

		for i := 0; i < len(*wallets); i++ {

			model = (*wallets)[i]

			if model.ID == walletId {

				fmt.Printf("\033[1A\033[KConfirm that is \033[1mWallet (%d) %s\033[0m? Type Y or N: ", model.ID, model.Title)
				fmt.Scanf("%s", &answer)
				break
			}
		}

		if answer == "Y" {
			fmt.Printf("\033[1A\033[KWallet: \033[1m(%d) %s.\033[0m", model.ID, model.Title)
		}

	}

	answer = ""
	var sourceId uint

	for answer != "Y" {
		if sourceId > 0 {
			PrintBold("\033[1A\033[KWhat is source id? ")
		} else {
			PrintBold("\nWhat is source id? ")
		}

		fmt.Scanf("%d", &sourceId)

		var model database.Source

		for i := 0; i < len(*sources); i++ {

			model = (*sources)[i]

			if model.ID == sourceId {

				fmt.Printf("\033[1A\033[KConfirm that is Source \033[1m(%d) %s (%s)\033[0m? Type Y or N: ", model.ID, model.Title, model.Flux)
				fmt.Scanf("%s", &answer)
				break
			}
		}

		if answer == "Y" {
			fmt.Printf("\033[1A\033[KSource: \033[1m(%d) %s (%s).\033[0m", model.ID, model.Title, model.Flux)
		}
	}

	//description
	reader := bufio.NewReader(os.Stdin)

	PrintBold("\nWhat is the description? ")
	text, _ := reader.ReadString('\n')

	description := strings.Split(text, "\n")[0]
	fmt.Printf("\033[1A\033[KDescription: \033[1m%s.\033[0m", description)

	//cash
	var cash float32

	PrintBold("\nHow much money? ")
	fmt.Scanf("%f", &cash)
	fmt.Printf("\033[1A\033[KCash: \033[1m%.2f.\033[0m", cash)

	fmt.Printf("\n\nConfirm flow insertion? Y or N: ")
	fmt.Scanf("%s", &answer)

	if answer == "Y" {

		if err := database.InsertFlow(
			&database.Flow{
				CategoryID:  categoryId,
				WalletID:    walletId,
				SourceID:    sourceId,
				Description: description,
				Cash:        cash,
			}); err != nil {
			fmt.Println(err)
		} else {
			PrintSuccess("\nNew flow added successfuly.\n")
		}
	}

}
