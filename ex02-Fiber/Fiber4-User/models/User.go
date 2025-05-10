package models

import (
	"time"

	"gorm.io/gorm"
)

// User model ORM
type User struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"required"`
	Email     string         `gorm:"unique"`
	Password  string         `gorm:"required"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
