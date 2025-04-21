package repository

import (
	"sync"
)

var (
	repositoryInit     sync.Once
	repositoryInstance *Repository
)

type Repository struct {
	// mongoDB (Access DB)

	User *UserRepository
}

func NewRepository() *Repository {
	repositoryInit.Do(func() {
		repositoryInstance = &Repository{
			User: newUserRepository(),
		}
	})

	return repositoryInstance
}
