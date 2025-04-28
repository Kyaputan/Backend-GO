package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	s3Bucket := os.Getenv("HELLO")
	fmt.Println(s3Bucket)
}
