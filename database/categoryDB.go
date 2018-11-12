package database

import "github.com/gabrielroriz/cli-fineasy/models"

func GetCategories() *[]models.Category {

	if values, ok := (*db).DB.
		Find(&[]models.Category{}).
		Value.(*[]models.Category); ok {

		return values
	}

	return nil
}
