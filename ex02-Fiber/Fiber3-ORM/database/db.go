package database

import (
	"log"

	"github.com/Kyaputan/Golang/ex02-Fiber/Fiber3-ORM/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// กำหนดตัวแปร db เพื่อใช้งานภายใน ใส่ ๆ มา
var DB *gorm.DB

// ฟังก์ชันการเชื่อมต่อฐานข้อมูล
func Connect() (*gorm.DB, error) {
	conn := config.GetDatabaseURL()
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	log.Println("✅ Connected to MySQL database")
	return db, nil
}

// ฟังก์ชันปิดการเชื่อมต่อฐานข้อมูล
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
