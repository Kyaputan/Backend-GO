package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	s3Bucket := os.Getenv("HELLO")

	router := gin.Default()


	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": s3Bucket})
	})

	router.Run(":5000")
}
