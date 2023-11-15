package controllers

import (
	"mksc_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CupInput struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}

func AddCup(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	var input CupInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cup := models.Cup{
		Name:  input.Name,
		Image: input.Image,
	}

	savedCup, err := cup.SaveCup()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"cup": savedCup})
}

func GetAllCups(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	cups := models.FindAllCups()

	context.JSON(http.StatusAccepted, cups)
}

func GetCupByID(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	cid, err := strconv.ParseUint(context.Param("cupID"), 10, 32)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cup, err := models.FindCupByID(uint(cid))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, cup)
}
