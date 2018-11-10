package models

import (
	"github.com/jinzhu/gorm"
)

type Source struct {
	gorm.Model
	Title string `gorm:"type:varchar(30);not null"`
	Flux  string `gorm:"not null; type:fluxType"`
}
