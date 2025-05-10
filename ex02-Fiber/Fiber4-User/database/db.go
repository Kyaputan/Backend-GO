package database

import (
	"log"

	"github.com/Kyaputan/Backend-GO/ex02-Fiber/Fiber4-User/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// กำหนดตัวแปร db เพื่อใช้งานภายใน ใส่ ๆ มา
var DB *gorm.DB

// ฟังก์ชันการเชื่อมต่อฐานข้อมูล
func Connect() {
	conn := config.GetDatabaseURL()
	var err error
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	log.Println("✅ Connected to MySQL database")
}

// ฟังก์ชันปิดการเชื่อมต่อฐานข้อมูล
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("❌ Error closing database:", err)
	}
	sqlDB.Close()
}
