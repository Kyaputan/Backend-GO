package models

import (
	"gorm.io/gorm"
)

// User model ORM
type Products struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:text"`
	Price int16  `gorm:"required"`
}
