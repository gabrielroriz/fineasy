package main

import (
	"fmt"

	db "github.com/gabrielroriz/cli-fineasy/database"
	"github.com/gabrielroriz/cli-fineasy/handlers"
)

func main() {

	db := db.InitDB()
	defer db.Close()

	command := ""
	for command != "\\q" {

		fmt.Print("fineasy>")
		fmt.Scanf("%s", &command)

		switch command {
		case "lf":
			handlers.ListFlows()
		default:

		}

	}
}
