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

	// values := database.GetFlows()

	// var total float32

	// fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n")
	// fmt.Printf(" %-10s | %-20s | %-20s | %-20s | %-20s | %-20s |\n", "id", "date", "source", "category", "wallet", "cash")
	// fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n")
	// for i := 0; i < len(*values); i++ {

	// 	model := (*values)[i]

	// 	total += model.Cash

	// 	date := fmt.Sprintf("%d/%d/%d", model.CreatedAt.Day(), model.CreatedAt.Month(), model.CreatedAt.Year())
	// 	source := fmt.Sprintf("(%d) %s", model.Source.ID, model.Source.Title)
	// 	category := fmt.Sprintf("(%d) %s", model.Category.ID, model.Category.Title)
	// 	wallet := fmt.Sprintf("(%d) %s", model.Wallet.ID, model.Wallet.Title)
	// 	cash := fmt.Sprintf("R$ %.2f", model.Cash)

	// 	fmt.Printf(" %-*s | %-*s | %-*s | %-*s | %-*s | %-*s |\n",
	// 		10, fmt.Sprint(model.ID),
	// 		20, date,
	// 		20, source,
	// 		20, category,
	// 		20, wallet,
	// 		20, cash)
	// }
	// fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n")
	// fmt.Printf("                                                                                                        | %-*s |\n", 20, fmt.Sprintf("R$ %.2f", total))
	// fmt.Printf("--------------------------------------------------------------------------------------------------------------------------------\n\n")
}
