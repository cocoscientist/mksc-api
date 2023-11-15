package main

import (
	"fmt"
	"log"
	"mksc_api/controllers"
	"mksc_api/database"
	"mksc_api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Character{})
	database.Database.AutoMigrate(&models.Cup{})
	database.Database.AutoMigrate(&models.Track{})
	database.Database.AutoMigrate(&models.Item{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)

	mainRoutes := router.Group("/api")
	mainRoutes.POST("/characters", controllers.AddCharacter)
	mainRoutes.GET("/character", controllers.GetCharacterByName)
	mainRoutes.GET("/characters", controllers.GetAllCharacters)
	mainRoutes.POST("/cups", controllers.AddCup)
	mainRoutes.GET("/cups", controllers.GetAllCups)
	mainRoutes.GET("/cup/:cupID", controllers.GetCupByID)
	mainRoutes.POST("/tracks", controllers.AddTrack)
	mainRoutes.GET("/tracks", controllers.GetAllTracks)
	mainRoutes.GET("/tracks/:cupID", controllers.GetTrackByCupID)
	mainRoutes.POST("/items", controllers.AddItem)
	mainRoutes.GET("/item", controllers.GetItemByName)
	mainRoutes.GET("/item/:itemID", controllers.GetItemByID)
	mainRoutes.GET("/items", controllers.GetAllItems)

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
