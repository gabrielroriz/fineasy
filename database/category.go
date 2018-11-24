package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Title string `sql:"not null; unique; type: varchar(30);"`
}

func (c Category) ToString() string {
	return fmt.Sprintf("(%d) %s", c.ID, c.Title)
}

func (c Category) GetID() uint {
	return c.ID
}

func (c Category) GetTypeInString() string {
	return "Category"
}

func GetCategories() []Category {

	if values, ok := (*dbConfig).DB.
		Find(&[]Category{}).
		Value.(*[]Category); ok {

		return *values
	}

	return nil
}

func GetCategories2() []Category {

	if values, ok := (*dbConfig).DB.
		Find(&[]Category{}).
		Value.([]Category); ok {

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
