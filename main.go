package main

import (
	"fmt"

	"github.com/gabrielroriz/fineasy/database"
	"github.com/gabrielroriz/fineasy/handlers"
)

func main() {

	var db *database.DBConfig
	var err error

	for db == nil {
		db, err = database.InitDB()
		if err != nil {
			fmt.Println(err)
			database.SetDBConfig(handlers.InsertDBConfig())
		} else {
			defer db.DB.Close()
		}
	}

	command := ""
	for command != "\\q" {

		fmt.Print("fineasy>")
		fmt.Scanf("%s", &command)

		switch command {
		case "lf":
			handlers.ListFlows()

		case "lc":
			handlers.ListCategories()

		case "lw":
			handlers.ListWallets()

		case "ls":
			handlers.ListSources()

		default:

		}

	}
}
