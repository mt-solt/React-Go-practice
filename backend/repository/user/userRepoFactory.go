package user

import "react-go-practice/db"

type UserRepoFactory interface {
	Create() UserRepository
}

type userRepoFactory struct {
	querier db.Querier
}

func New(querier db.Querier) UserRepoFactory {
	return &userRepoFactory{
		querier: querier,
	}
}

func (f userRepoFactory) Create() UserRepository {
	return NewUserRepoPostgres(f.querier)
}
