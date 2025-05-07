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

	PORT := os.Getenv("PORT")
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.JSON(fiber.Map{
				"message": "Hello " + c.Params("name"),
				"status":  "success",
			})
		}
		return c.JSON(fiber.Map{
			"message": "✅ Server is running! Fiber is ready to go!",
			"status":  "success",
		})
	})

	app.Post("/post/:name?", func(c *fiber.Ctx) error {
		var json map[string]interface{}
		if err := c.BodyParser(&json); err != nil {
			fmt.Println("Error /post/:name?")
			return c.JSON(fiber.Map{
				"message": "❌ Failed to bind JSON",
				"status":  "error",
			})
		}
		if c.Params("name") != "" {
			return c.JSON(fiber.Map{
				"message": "✅ JSON received successfully!",
				"status":  "success",
				"AllData": map[string]interface{}{"json": json,
					"name": c.Params("name")},
			})
		}
		return c.JSON(fiber.Map{
			"message": "✅ JSON received successfully!",
			"status":  "success",
			"AllData": json,
		})
	})

	if err := app.Listen(":" + PORT); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}

}
