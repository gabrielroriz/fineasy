package database

import "github.com/gabrielroriz/cli-fineasy/models"

//GetFlows : Get all flow tuples
func GetFlows() *[]models.Flow {

	if values, ok := (*db).DB.
		Preload("Source").
		Preload("Category").
		Preload("Wallet").
		Find(&[]models.Flow{}).
		Value.(*[]models.Flow); ok {

		return values
	}

	return nil
}
