package repository

import (
	"gorm.io/gorm"
)

// Struct
type Repository struct {
	DB *gorm.DB
}