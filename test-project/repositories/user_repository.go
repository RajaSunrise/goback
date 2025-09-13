package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/test/test-project/models"
	"github.com/google/uuid"
)

// Tipe ID dinamis berdasarkan database
type UserID = uuid.UUID

// UserRepository adalah interface konsisten untuk operasi data User.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindAll(ctx context.Context, page, perPage int) ([]models.User, int64, error)
	FindByID(ctx context.Context, id UserID) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id UserID) error
}
