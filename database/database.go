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

		db.DB, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=fineasy password=%s sslmode=disable",
				db.Host, db.Port, db.User, db.Password))

		if err != nil {
			return nil, err
		}

		db.DB.AutoMigrate(
			&Source{},
			&Wallet{},
			&Category{},
			&Flow{})

		// db.DB.LogMode(true)

		return db, nil
	}
}
