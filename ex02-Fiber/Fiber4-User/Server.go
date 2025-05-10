package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kyaputan/Backend-GO/ex02-Fiber/Fiber4-User/config"
	"github.com/Kyaputan/Backend-GO/ex02-Fiber/Fiber4-User/database"
	"github.com/Kyaputan/Backend-GO/ex02-Fiber/Fiber4-User/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type SignupInput struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type LoginInput struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func main() {

	database.Connect()
	defer database.Close()
	app := fiber.New()

	PORT := os.Getenv("PORT")

	fmt.Println(config.GetDatabaseURL())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ Server is running! Fiber is ready to go!")
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User
		result := database.DB.Find(&users)

		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"users":  users,
			"status": "success",
		})
	})

	app.Get("/users/:ID?", func(c *fiber.Ctx) error {
		ID := c.Params("ID")
		if ID != "" {
			var user models.User
			result := database.DB.First(&user, "ID = ?", ID)

			if result.Error != nil {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": result.Error.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"user":   user,
				"status": "success",
			})
		}

		return c.JSON(fiber.Map{
			"message": "✅ Server is running! Fiber is ready to go!",
			"status":  "success",
		})
	})

	app.Post("/users/singup", func(c *fiber.Ctx) error {

		var input SignupInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if input.Name == "" || input.Email == "" || input.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "All fields are required",
			})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}

		user := models.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: string(hashedPassword),
		}

		if err := database.DB.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Could not create user: " + err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created successfully",
			"user": fiber.Map{
				"ID":    user.ID,
				"Name":  user.Name,
				"Email": user.Email,
			},
		})
	})

	app.Post("/users/login", func(c *fiber.Ctx) error {
		var input LoginInput
		var user models.User

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if input.Email == "" || input.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "All fields are required",
			})
		}

		result := database.DB.Where("email = ?", input.Email).First(&user)

		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid password",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Login successful",
			"user": fiber.Map{
				"ID":    user.ID,
				"Name":  user.Name,
				"Email": user.Email,
			},
		})
	})

	if err := app.Listen(":" + PORT); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
