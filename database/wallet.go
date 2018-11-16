package database

import (
	"github.com/jinzhu/gorm"
)

type Wallet struct {
	gorm.Model
	Title string `sql:"not null; type: varchar(30);"`
}

func GetWallets() *[]Wallet {

	if values, ok := (*db).DB.
		Find(&[]Wallet{}).
		Value.(*[]Wallet); ok {

		return values
	}

	return nil
}
