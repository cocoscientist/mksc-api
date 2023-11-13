package controllers

import (
	"mksc_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CharacterInput struct {
	Name   string `json:"name" binding:"required"`
	Image  string `json:"image" binding:"required"`
	Speed  int    `json:"speed" binding:"required"`
	Weight int    `json:"weight" binding:"required"`
}

func AddCharacter(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	var input CharacterInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	character := models.Character{
		Name:   input.Name,
		Image:  input.Image,
		Speed:  input.Speed,
		Weight: input.Weight,
	}

	savedCharacter, err := character.SaveCharacter()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"character": savedCharacter})
}

func GetCharacterByName(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	charName := context.Query("name")
	chara, err := models.FindCharacterByName(charName)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"result": chara})
}
