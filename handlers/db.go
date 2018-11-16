package handlers

import (
	"fmt"

	"github.com/gabrielroriz/fineasy/database"
)

func InsertDBConfig() *database.DBConfig {

	dbConfig := &database.DBConfig{}

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
