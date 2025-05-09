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
	// โหลดค่าจากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("❌ Error loading .env file")
		return
	}

	// ดึงค่าพอร์ตจาก ENV
	PORT := os.Getenv("PORT")

	// สร้างแอป Fiber ใหม่
	app := fiber.New()

	// เปิดใช้งาน middleware สำหรับจัดการ CORS
	app.Use(cors.New())

	// ----------------------------
	// Route: GET /:name?
	// แสดงข้อความต้อนรับ (รองรับชื่อแบบ optional)
	// ----------------------------
	app.Get("/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name != "" {
			return c.JSON(fiber.Map{
				"message": "Hello " + name,
				"status":  "success",
			})
		}

		return c.JSON(fiber.Map{
			"message": "✅ Server is running! Fiber is ready to go!",
			"status":  "success",
		})
	})

	// ----------------------------
	// Route: POST /post/:name?
	// รับข้อมูล JSON จาก client และส่งกลับพร้อมชื่อ (ถ้ามี)
	// ----------------------------
	type RequestBody struct {
		Email string `json:"email"`
		Age   int    `json:"age"`
	}
	app.Post("/post/:name?", func(c *fiber.Ctx) error {
		var json RequestBody
		name := c.Params("name")

		// แปลง body ที่รับมาเป็น JSON
		if err := c.BodyParser(&json); err != nil {
			fmt.Println("❌ Error parsing JSON at /post/:name?")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "❌ Failed to bind JSON",
				"status":  "error",
			})
		}

		if name != "" {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "✅ JSON received successfully!",
				"status":  "success",
				"data": map[string]any{
					"json": json,
					"Name": name,
				},
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "✅ JSON received successfully!",
			"status":  "success",
			"data":    json,
		})
	})

	// ----------------------------
	// เริ่มต้นเซิร์ฟเวอร์
	// ----------------------------
	if err := app.Listen(":" + PORT); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
