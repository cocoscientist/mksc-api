package controllers

import (
	"mksc_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemInput struct {
	Name        string `json:"name" binding:"required"`
	Icon        string `json:"icon" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func AddItem(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	var input ItemInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{
		Name:        input.Name,
		Icon:        input.Icon,
		Description: input.Description,
	}

	savedItem, err := item.SaveItem()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"item": savedItem})
}

func GetAllItems(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	items := models.GetAllItems()

	context.JSON(http.StatusAccepted, items)
}

func GetItemByName(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	itemName := context.Query("name")
	item, err := models.FindItemByName(itemName)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, item)
}

func GetItemByID(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	iid, err := strconv.ParseUint(context.Param("itemID"), 10, 32)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := models.GetItemByID(uint(iid))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, item)
}
