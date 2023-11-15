package models

import (
	"mksc_api/database"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Icon        string `gorm:"not null;unique" json:"icon"`
	Description string `gorm:"type:text" json:"description"`
}

func (item *Item) SaveItem() (*Item, error) {
	err := database.Database.Create(&item).Error
	if err != nil {
		return &Item{}, err
	}
	return item, nil
}

func GetAllItems() []Item {
	var items []Item
	result := database.Database.Find(&items)
	if result.Error != nil {
		return nil
	}
	return items
}

func GetItemByID(id uint) (Item, error) {
	var curItem Item
	erro := database.Database.Where("ID=?", id).Find(&curItem).Error
	if erro != nil {
		return Item{}, erro
	}
	return curItem, nil
}

func FindItemByName(name string) (Item, error) {
	var item Item
	erro := database.Database.Where("name=?", name).Find(&item).Error
	if erro != nil {
		return Item{}, erro
	}
	return item, nil
}
