package models

import "github.com/jinzhu/gorm"

type Flow struct {
	gorm.Model
	Cash       float32 `gorm:"numeric(8,2);"`
	Flux       string  `gorm:"not null;type:fluxType;"`
	SourceID   uint
	WalletID   uint
	CategoryID uint

	Source   Source   `gorm:"foreignkey:SourceID"`
	Wallet   Wallet   `gorm:"foreignkey:WalletID"`
	Category Category `gorm:"foreignkey:CategoryID"`
}
