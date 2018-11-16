package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//InitDB : Init DB
func InitDB() (*DBConfig, error) {

	var err error

	err = readDBConfig()

	if err != nil {
		return nil, err
	} else {

		dbConfig.DB, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Database, dbConfig.Password))

		dbConfig.DB.LogMode(false)

		if err != nil {
			return nil, err
		}

		dbConfig.DB.AutoMigrate(
			&Source{},
			&Wallet{},
			&Category{},
			&Flow{})

		dbConfig.DB.Exec("CREATE TYPE fluxType as enum('expense', 'income');")

		return dbConfig, nil
	}
}
