package handlers

import (
	"fmt"

	db "github.com/gabrielroriz/cli-fineasy/database"
)

func ListWallets() {

	values := db.GetWallets()

	fmt.Printf("------------------------------------\n")
	fmt.Printf(" %-10s | %-20s |\n", "id", "title")
	fmt.Printf("------------------------------------\n")
	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		fmt.Printf(" %-*s | %-*s |\n",
			10, fmt.Sprint(model.ID),
			20, model.Title)
	}
	fmt.Printf("------------------------------------\n")

}
