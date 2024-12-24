package main

import (
	"fmt"

	"fineasy/database"
	"fineasy/handlers"
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

		fmt.Print("\nfineasy> ")
		fmt.Scanf("%s", &command)

		switch command {

		//lists
		case "lf":
			handlers.ListFlows()

		case "lc":
			handlers.ListCategories()

		case "lw":
			handlers.ListWallets()

		case "ls":
			handlers.ListSources()

		//insertions
		case "if":
			handlers.InsertFlow()

		case "ic":
			handlers.InsertCategory()

		case "iw":
			handlers.InsertWallet()

		case "is":
			handlers.InsertSource()

		default:

		}

	}
}
