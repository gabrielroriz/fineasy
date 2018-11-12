package handlers

import (
	"fmt"

	db "github.com/gabrielroriz/cli-fineasy/database"
)

func ListSources() {

	values := db.GetSources()

	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf(" %-10s | %-20s | %-20s |\n", "id", "title", "flux")
	fmt.Printf("-----------------------------------------------------------\n")
	for i := 0; i < len(*values); i++ {

		model := (*values)[i]

		fmt.Printf(" %-*s | %-*s | %-*s |\n",
			10, fmt.Sprint(model.ID),
			20, model.Title,
			20, model.Flux)
	}
	fmt.Printf("-----------------------------------------------------------\n")

}
