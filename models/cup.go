package models

import (
	"mksc_api/database"

	"gorm.io/gorm"
)

type Cup struct {
	gorm.Model
	Name   string  `gorm:"size:255;not null;unique" json:"name"`
	Image  string  `gorm:"not null;unique" json:"image"`
	Tracks []Track `json:"-"`
}

func (cup *Cup) SaveCup() (*Cup, error) {
	err := database.Database.Create(&cup).Error
	if err != nil {
		return &Cup{}, err
	}
	return cup, nil
}

func FindCupByName(name string) (Cup, error) {
	var curCup Cup
	erro := database.Database.Where("name=?", name).Find(&curCup).Error
	if erro != nil {
		return Cup{}, erro
	}
	return curCup, nil
}

func FindCupByID(id uint) (Cup, error) {
	var curCup Cup
	erro := database.Database.Where("ID=?", id).Find(&curCup).Error
	if erro != nil {
		return Cup{}, erro
	}
	return curCup, nil
}

func FindAllCups() []Cup {
	var cups []Cup
	result := database.Database.Find(&cups)
	if result.Error != nil {
		return nil
	}
	return cups
}
