// internal/models/base_model.go

package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel defines the common fields for GORM models.
type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// You can add other base models or model-related utility functions here.

// User represents a user in the system.
type User struct {
	BaseModel
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}
