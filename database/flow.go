package database

import (
	"fmt"

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

// GetFlows : Get all flow tuples
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

func (f Flow) ToTableFormat() []string {
	id := fmt.Sprintf("%d", f.ID)
	date := fmt.Sprintf("%d/%d/%d", f.CreatedAt.Day(), f.CreatedAt.Month(), f.CreatedAt.Year())
	source := fmt.Sprintf("(%d) %s", f.Source.ID, f.Source.Title)
	flux := fmt.Sprintf("%s", f.Source.Flux)
	description := fmt.Sprintf("%s", f.Description)
	category := fmt.Sprintf("(%d) %s", f.Category.ID, f.Category.Title)
	wallet := fmt.Sprintf("(%d) %s", f.Wallet.ID, f.Wallet.Title)
	cash := fmt.Sprintf("R$ %.2f", f.Cash)
	return []string{id, date, source, flux, description, category, wallet, cash}
}

func (f Flow) InMemoryTableFormat() []string {

	categoryId := fmt.Sprintf("%d", f.CategoryID)
	walletId := fmt.Sprintf("%d", f.WalletID)
	sourceId := fmt.Sprintf("%d", f.SourceID)
	description := fmt.Sprintf("%s", f.Description)
	cash := fmt.Sprintf("R$ %.2f", f.Cash)

	return []string{categoryId, walletId, sourceId, description, cash}
}
