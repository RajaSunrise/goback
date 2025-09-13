package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User adalah model data utama.
// Struct tags disesuaikan berdasarkan pilihan ORM.
type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// CreateUserRequest adalah DTO (Data Transfer Object) untuk membuat user.
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UpdateUserRequest adalah DTO untuk memperbarui user.
type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	IsActive *bool   `json:"is_active,omitempty"`
}

// UserResponse adalah DTO untuk menampilkan data user (tanpa password).
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Pagination adalah struct untuk informasi paginasi.
type Pagination struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	TotalPages int   `json:"total_pages"`
}

// UsersResponse adalah DTO untuk daftar user dengan paginasi.
type UsersResponse struct {
	Users      []UserResponse `json:"users"`
	Pagination Pagination     `json:"pagination"`
}