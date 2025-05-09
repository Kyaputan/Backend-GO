package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// โหลดไฟล์ .env เฉย ๆ
func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}
}

// root:310146@tcp(127.0.0.1:3306)/My_SQL?charset=utf8mb4&parseTime=True&loc=Local
func GetDatabaseURL() string {
	LoadEnvVariables()
	dbURL := os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASS") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dbURL
}
