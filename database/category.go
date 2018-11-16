package database

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Title string `sql:"not null; unique; type: varchar(30);"`
}

func GetCategories() *[]Category {

	if values, ok := (*dbConfig).DB.
		Find(&[]Category{}).
		Value.(*[]Category); ok {

		return values
	}

	return nil
}

func InsertCategory(category *Category) error {

	if err := (*dbConfig).DB.Create(category).Error; err != nil {
		return err
	}

	return nil
}
