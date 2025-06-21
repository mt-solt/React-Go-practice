package user

import (
	"context"

	"react-go-practice/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user models.CreateUserRequest) (*models.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	UpdatePassword(ctx context.Context, id uuid.UUID, passwordHash string) (*models.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
