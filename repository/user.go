package repository

import "github.com/BEpaul/go-test-server/types"

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{}, // 초기화 진행해야 NPE 뜨지 않음
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Update(beforeUser *types.User, updatedUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
