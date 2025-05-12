package models

import (
	"gorm.io/gorm"
)

// User model ORM
type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"required"`
	Email    string `gorm:"unique"`
	Password string `gorm:"required"`
	Role     string `gorm:"default:'user'"`
}
