package handlers

import (
	"fmt"

	"fineasy/database"
)

func InsertDBConfig() *database.DBConfig {

	dbConfig := database.DBConfig{}

	fmt.Printf("host: ")
	fmt.Scanf("%s", &(dbConfig.Host))

	fmt.Printf("port: ")
	fmt.Scanf("%s", &(dbConfig.Port))

	fmt.Printf("database: ")
	fmt.Scanf("%s", &(dbConfig.Database))

	fmt.Printf("username: ")
	fmt.Scanf("%s", &(dbConfig.User))

	fmt.Printf("password: ")
	fmt.Scanf("%s", &(dbConfig.Password))

	return &dbConfig
}
