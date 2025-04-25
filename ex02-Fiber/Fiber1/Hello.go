package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Access the environment variable
	envget := os.Getenv("HELLO")
	fmt.Printf("\nText : %s\n", envget)
}
