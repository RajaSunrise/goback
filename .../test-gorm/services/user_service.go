package services

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/user/test-gorm/internal/models"
	"github.com/user/test-gorm/internal/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

// Tipe ID dinamis
type UserID = uuid.UUID

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error)
	GetAllUsers(ctx context.Context, page, perPage int) (*models.UsersResponse, error)
	GetUserByID(ctx context.Context, id UserID) (*models.UserResponse, error)
	UpdateUser(ctx context.Context, id UserID, req *models.UpdateUserRequest) (*models.UserResponse, error)
	DeleteUser(ctx context.Context, id UserID) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error) {
	existingUser, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Active: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) GetUserByID(ctx context.Context, id UserID) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toUserResponse(user), nil
}

func (s *userService) GetAllUsers(ctx context.Context, page, perPage int) (*models.UsersResponse, error) {
	users, total, err := s.userRepo.FindAll(ctx, page, perPage)
	if err != nil {
		return nil, err
	}

	userResponses := make([]models.UserResponse, 0, len(users))
	for _, user := range users {
		userResponses = append(userResponses, *toUserResponse(&user))
	}

	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &models.UsersResponse{
		Users: userResponses,
		Pagination: models.Pagination{
			Total:      total,
			Page:       page,
			PerPage:    perPage,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, id UserID, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Active != nil {
		user.Active = *req.Active
	}
	user.UpdatedAt = time.Now()

	user, err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) DeleteUser(ctx context.Context, id UserID) error {
	return s.userRepo.Delete(ctx, id)
}


// toUserResponse adalah helper untuk mengonversi model User ke DTO UserResponse.
func toUserResponse(u *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Active:  u.Active,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
