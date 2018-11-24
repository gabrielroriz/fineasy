package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Wallet struct {
	gorm.Model
	Title string `sql:"not null; type: varchar(30);"`
}

func (w Wallet) ToString() string {
	return fmt.Sprintf("(%d) %s", w.ID, w.Title)
}

func (w Wallet) GetID() uint {
	return w.ID
}

func (w Wallet) GetTypeInString() string {
	return "Wallet"
}

func GetWallets() []Wallet {

	if values, ok := (*dbConfig).DB.
		Find(&[]Wallet{}).
		Value.(*[]Wallet); ok {

		return *values
	}

	return nil
}

func InsertWallet(wallet *Wallet) error {

	if err := (*dbConfig).DB.Create(wallet).Error; err != nil {
		return err
	}

	return nil
}
