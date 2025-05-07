package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	HELLO := os.Getenv("HELLO")

	app := fiber.New()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/" {
			return c.Next()
		}
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": HELLO})
	})

	log.Fatal(app.Listen(":5000"))
}
