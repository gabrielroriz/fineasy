package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Title string `sql:"not null; unique; type: varchar(30);"`
}
