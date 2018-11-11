package main

import (
	"os"

	db "github.com/gabrielroriz/cli-fineasy/database"
	"github.com/gabrielroriz/cli-fineasy/handlers"
)

func main() {

	db.InitDB()

	args := os.Args[1:]

	if len(args) != 0 && args[0] != "" {
		switch args[0] {
		case "-lf":
			handlers.ListFlows()
		default:

		}
	}
}
