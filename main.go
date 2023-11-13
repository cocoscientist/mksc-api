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

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
