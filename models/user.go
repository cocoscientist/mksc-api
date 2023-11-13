package models

import (
	"html"
	"mksc_api/database"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	ApiKey   string `gorm:"size:255;not null;unique;" json:"-"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	apiKeyHash, err := bcrypt.GenerateFromPassword([]byte(user.ApiKey), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.ApiKey = string(apiKeyHash)
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func getBcryptedKey(apiKey string) (string, error) {
	apiKeyH, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(apiKeyH), nil
}

func FindUserByApiKey(apiKey string) (User, error) {
	apiKeyHash, err := getBcryptedKey(apiKey)
	if err != nil {
		return User{}, err
	}
	var user User
	erro := database.Database.Where("api_key=?", apiKeyHash).Find(&user).Error
	if erro != nil {
		return User{}, erro
	}
	return user, nil
}
