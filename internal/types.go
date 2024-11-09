package internal

import (
	"playground/internal/entity"
)

type Repository interface {
	GetUsers(ids []int64) ([]*entity.User, error)
	CreateUsers(users []*entity.User) error
}

type RepositoryConfig interface{}

type TokenAuthorization interface {
}
