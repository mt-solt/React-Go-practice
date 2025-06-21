package user

import (
	"context"
	"react-go-practice/db"
	"react-go-practice/models"

	"github.com/google/uuid"
)

type postgresRepo struct {
	querier db.Querier
}

func NewUserRepoPostgres(querier db.Querier) UserRepository {
	return &postgresRepo{
		querier: querier,
	}
}

func (r *postgresRepo) Create(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	params := db.CreateUserParams{
		Username:     req.Username,
		PasswordHash: req.Password,
	}
	dbUser, err := r.querier.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return convertDBUserToModel(&dbUser), nil
}

func (r *postgresRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	dbUser, err := r.querier.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return convertDBUserToModel(&dbUser), nil
}

func (r *postgresRepo) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	dbUser, err := r.querier.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return convertDBUserToModel(&dbUser), nil
}

func (r *postgresRepo) List(ctx context.Context) ([]*models.User, error) {
	dbUsers, err := r.querier.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = convertDBUserToModel(&dbUser)
	}
	return users, nil
}

func (r *postgresRepo) UpdatePassword(ctx context.Context, id uuid.UUID, passwordHash string) (*models.User, error) {
	params := db.UpdateUserPasswordParams{
		ID:           id,
		PasswordHash: passwordHash,
	}
	dbUser, err := r.querier.UpdateUserPassword(ctx, params)
	if err != nil {
		return nil, err
	}
	return convertDBUserToModel(&dbUser), nil
}

func (r *postgresRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.querier.DeleteUser(ctx, id)
}

func convertDBUserToModel(dbUser *db.User) *models.User {
	return &models.User{
		ID:           dbUser.ID,
		Username:     dbUser.Username,
		PasswordHash: dbUser.PasswordHash,
	}
}
