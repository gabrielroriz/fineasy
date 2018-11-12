package handlers

import (
	"fmt"

	db "github.com/gabrielroriz/cli-fineasy/database"
)

func ConfigDB() *db.DBConfig {

	dbConfig := &db.DBConfig{}

	fmt.Printf("host=")
	fmt.Scanf("%s", &((*dbConfig).Host))

	fmt.Printf("port=")
	fmt.Scanf("%s", &((*dbConfig).Port))

	fmt.Printf("user=")
	fmt.Scanf("%s", &((*dbConfig).User))

	fmt.Printf("password=")
	fmt.Scanf("%s", &((*dbConfig).Password))

	return dbConfig
}
