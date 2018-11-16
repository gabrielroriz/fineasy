package database

import (
	"github.com/jinzhu/gorm"
)

type Source struct {
	gorm.Model
	Title string `sql:"not null; type: varchar(30);"`
	Flux  string `sql:"not null; type: fluxType;"`
}

func GetSources() *[]Source {

	if values, ok := (*db).DB.
		Find(&[]Source{}).
		Value.(*[]Source); ok {

		return values
	}

	return nil
}
