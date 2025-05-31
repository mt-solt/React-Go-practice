package random

import (
	"context"

	"react-go-practice/models"
)

type RandomRepository interface {
	Create(ctx context.Context, data models.Random) error
	Read(ctx context.Context, id int) (models.Random, error)
	Update(ctx context.Context, id int, random int64) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.Random, error)
	QueryByUser(ctx context.Context, id int) ([]models.Random, error)
}
