package database

import "github.com/gabrielroriz/cli-fineasy/models"

func GetSources() *[]models.Source {

	if values, ok := (*db).DB.
		Find(&[]models.Source{}).
		Value.(*[]models.Source); ok {

		return values
	}

	return nil
}
