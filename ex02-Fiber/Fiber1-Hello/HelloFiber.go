package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	Hello := os.Getenv("HELLO")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(Hello)
	})

	app.Listen(":5000")
}
