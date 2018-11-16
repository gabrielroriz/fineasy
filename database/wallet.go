package database

import (
	"github.com/jinzhu/gorm"
)

type Wallet struct {
	gorm.Model
	Title string `sql:"not null; type: varchar(30);"`
}

func GetWallets() *[]Wallet {

	if values, ok := (*dbConfig).DB.
		Find(&[]Wallet{}).
		Value.(*[]Wallet); ok {

		return values
	}

	return nil
}

func InsertWallet(wallet *Wallet) error {

	if err := (*dbConfig).DB.Create(wallet).Error; err != nil {
		return err
	}

	return nil
}
