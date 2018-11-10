package models

import (
	"github.com/jinzhu/gorm"
)

type Flow struct {
	gorm.Model
	Cash       float32 `sql:"not null; type: numeric(8,2);"`
	Flux       string  `sql:"not null; type:fluxType;"`
	SourceID   uint    `sql:"type:int REFERENCES sources(id)"`
	WalletID   uint    `sql:"type:int REFERENCES wallets(id)"`
	CategoryID uint    `sql:"type:int REFERENCES categories(id)"`

	Source   Source   `gorm:"foreignkey:SourceID"`
	Wallet   Wallet   `gorm:"foreignkey:WalletID"`
	Category Category `gorm:"foreignkey:CategoryID"`
}
