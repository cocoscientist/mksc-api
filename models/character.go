package models

import (
	"mksc_api/database"

	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name   string `gorm:"size:255;not null;unique" json:"name"`
	Image  string `gorm:"not null;unique" json:"image"`
	Speed  int    `gorm:"not null" json:"speed"`
	Weight int    `gorm:"not null" json:"weight"`
}

func (character *Character) SaveCharacter() (*Character, error) {
	err := database.Database.Create(&character).Error
	if err != nil {
		return &Character{}, err
	}
	return character, nil
}

func FindCharacterByName(name string) (Character, error) {
	var chara Character
	erro := database.Database.Where("name=?", name).Find(&chara).Error
	if erro != nil {
		return Character{}, erro
	}
	return chara, nil
}
