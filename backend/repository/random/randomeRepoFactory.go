package random

type RandomRepoFactory interface {
	Create() RandomRepository
}

type randomRepoFactory struct {
}

func New() RandomRepoFactory {
	return &randomRepoFactory{}
}

func (f randomRepoFactory) Create() RandomRepository {
	return &postgresRepo{}
}
