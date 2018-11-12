package database

import "github.com/gabrielroriz/cli-fineasy/models"

func GetWallets() *[]models.Wallet {

	if values, ok := (*db).DB.
		Find(&[]models.Wallet{}).
		Value.(*[]models.Wallet); ok {

		return values
	}

	return nil
}
