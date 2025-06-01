package random

import (
	"context"

	"react-go-practice/models"
)

type RandomRepository interface {
	Create(ctx context.Context, data models.Random) error
	Read(ctx context.Context, uuid string) (models.Random, error)
	Update(ctx context.Context, uuid string, random int64) error
	Delete(ctx context.Context, uuid string) error
	GetAll(ctx context.Context) ([]models.Random, error)
	QueryByUser(ctx context.Context, userId string) ([]models.Random, error)
}
