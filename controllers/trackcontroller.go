package controllers

import (
	"mksc_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrackInput struct {
	Name  string `json:"name" binding:"required"`
	Icon  string `json:"icon" binding:"required"`
	Map   string `json:"map" binding:"required"`
	CupID uint   `json:"cupID" binding:"required"`
	Extra *bool  `json:"extra" binding:"required"`
}

func AddTrack(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	var input TrackInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{
		Name:  input.Name,
		Icon:  input.Icon,
		Map:   input.Map,
		CupID: input.CupID,
		Extra: *input.Extra,
	}

	savedTrack, err := track.SaveTrack()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"track": savedTrack})
}

func GetAllTracks(context *gin.Context) {
	token := context.Request.Header["Key"][0]

	user, erro := models.FindUserByApiKey(token)

	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	println(user.Username)

	tracks := models.FindAllTracks()

	context.JSON(http.StatusAccepted, tracks)
}

func GetTrackByCupID(context *gin.Context) {
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

	tracks, err := models.FindTrackByCupID(uint(cid))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, tracks)
}
