package controllers

import (
	"mksc_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiKeyGenerator(username string) (result string) {
	for _, v := range username {
		result = string(v) + result
	}
	return
}

func Register(context *gin.Context) {
	var input models.RegisterInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiBase := apiKeyGenerator(input.Username)

	user := models.User{
		Username: input.Username,
		Password: input.Password,
		ApiKey:   apiBase,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser, "apiKey": apiBase})
}
