package models

import "github.com/jinzhu/gorm"

type Wallet struct {
	gorm.Model
	Title string `sql:"not null; type: varchar(30);"`
}
