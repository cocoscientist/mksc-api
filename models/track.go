package models

import (
	"mksc_api/database"

	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	Name  string `gorm:"size:255;not null;unique" json:"name"`
	Icon  string `gorm:"not null;unique" json:"icon"`
	Map   string `gorm:"not null;unique" json:"map"`
	Extra bool   `gorm:"not null" json:"extra"`
	CupID uint
}

func (track *Track) SaveTrack() (*Track, error) {
	err := database.Database.Create(&track).Error
	if err != nil {
		return &Track{}, err
	}
	return track, nil
}

func FindTrackByName(name string) (Track, error) {
	var curTrack Track
	erro := database.Database.Where("name=?", name).Find(&curTrack).Error
	if erro != nil {
		return Track{}, erro
	}
	return curTrack, nil
}

func FindAllTracks() []Track {
	var tracks []Track
	result := database.Database.Find(&tracks)
	if result.Error != nil {
		return nil
	}
	return tracks
}

func FindTrackByCupID(cupID uint) ([]Track, error) {
	var tracks []Track
	erro := database.Database.Where("cup_ID=?", cupID).Find(&tracks).Error
	if erro != nil {
		return nil, erro
	}
	return tracks, nil
}
