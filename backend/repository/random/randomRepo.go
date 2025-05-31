package random

import (
	"context"
)

type RandomRepository interface {
	Create(ctx context.Context, random int64) error
	Read(ctx context.Context, id int) (int64, error)
	Update(ctx context.Context, id int, random int64) error
	Delete(ctx context.Context, id int) error
}

type postgresRepo struct {
}

func (r *postgresRepo) Create(ctx context.Context, random int64) error {
	return nil
}

func (r *postgresRepo) Read(ctx context.Context, id int) (int64, error) {
	return 0, nil
}

func (r *postgresRepo) Update(ctx context.Context, id int, random int64) error {
	return nil
}

func (r *postgresRepo) Delete(ctx context.Context, id int) error {
	return nil
}
