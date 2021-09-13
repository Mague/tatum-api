package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Mague/tatum-api/middlewares"
)

var router *gin.Engine

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	httpPort := os.Getenv("HTTP_PORT")
	//tatumApiUrl := os.Getenv("TATUM_API_URL")
	//tatumApiKey := os.Getenv("TATUM_API_KEY")
	gin.SetMode(gin.DebugMode)

	router = gin.Default()
	router.Use(middlewares.RateLimit, gin.Recovery())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Running",
		})
	})
	initializeRoutes()
	router.Run(":" + httpPort)
}
