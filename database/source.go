package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Source struct {
	gorm.Model
	Title string `sql:"not null; type: varchar(30);"`
	Flux  string `sql:"not null; type: fluxType;"`
}

func (s Source) ToString() string {
	return fmt.Sprintf("(%d) %s (%s)", s.ID, s.Title, s.Flux)
}

func (s Source) GetID() uint {
	return s.ID
}

func (s Source) GetTypeInString() string {
	return "Source"
}

func (s Source) ToTableFormat() []string {
	return []string{fmt.Sprintf("%d", s.ID), s.Title, s.Flux}
}

func GetSources() []Source {

	if values, ok := (*dbConfig).DB.
		Find(&[]Source{}).
		Value.(*[]Source); ok {

		return *values
	}

	return nil
}

func InsertSource(source *Source) error {

	if err := (*dbConfig).DB.Create(source).Error; err != nil {
		return err
	}

	return nil
}
