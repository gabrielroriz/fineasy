package database

import (
	"github.com/jinzhu/gorm"
)

type Flow struct {
	gorm.Model

	Cash        float32 `sql:"not null; type: numeric(8,2);"`
	Description string  `sql:"not null; type: varchar(50);"`
	SourceID    uint    `sql:"type:int REFERENCES sources(id)"`
	WalletID    uint    `sql:"type:int REFERENCES wallets(id)"`
	CategoryID  uint    `sql:"type:int REFERENCES categories(id)"`

	Source   Source   `gorm:"foreignkey:SourceID;"`
	Wallet   Wallet   `gorm:"foreignkey:WalletID;"`
	Category Category `gorm:"foreignkey:CategoryID;"`
}

//GetFlows : Get all flow tuples
func GetFlows() *[]Flow {

	if values, ok := (*dbConfig).DB.
		Preload("Source").
		Preload("Category").
		Preload("Wallet").
		Find(&[]Flow{}).
		Value.(*[]Flow); ok {

		return values
	}

	return nil
}

func InsertFlow(flow *Flow) error {

	if err := (*dbConfig).DB.Create(flow).Error; err != nil {
		return err
	}

	return nil
}
