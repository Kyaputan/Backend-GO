package main

import (
	"fmt"
	"log"

	"github.com/Kyaputan/Backend-GO/ex02-Fiber/Fiber3-ORM/config"
	"github.com/Kyaputan/Backend-GO/ex02-Fiber/Fiber3-ORM/database"
	"github.com/Kyaputan/Backend-GO/ex02-Fiber/Fiber3-ORM/models"
)

func main() {

	database.Connect()
	defer database.Close()

	fmt.Println(config.GetDatabaseURL())
	if err := database.DB.AutoMigrate(&models.Products{}); err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	newProduct := models.Products{
		Name:  "Pencil",
		Price: 5000,
	}
	result := database.DB.Create(&newProduct)

	if result.Error != nil {
		log.Fatal("❌ Insert failed:", result.Error)
	}

	fmt.Println("✅ Product created with ID:", newProduct.ID)

	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	}

	newUser := models.User{
		Name:     "Singkhet2",
		Email:    "Singkhet2@example.com",
		Password: "password",
	}
	result = database.DB.Create(&newUser)

	if result.Error != nil {
		log.Fatal("❌ Insert failed:", result.Error)
	}

	fmt.Println("✅ User created with ID:", newUser.ID)

}
