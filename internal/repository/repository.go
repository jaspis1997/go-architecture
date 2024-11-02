package repository

import (
	"context"
	"errors"
	"playground"
)

var (
	ErrorUnsupportedRepository = errors.New("unsupported repository")
)

type Repository struct {
	db Database
}

var migrateEntities []any

func RegisterEntities(entities []any) {
	migrateEntities = entities
}

var migrate func(db Database, targets []any) error

func RegisterMigrate(f func(db Database, targets []any) error) {
	migrate = f
}

func Migrate(repository playground.RepositoryConfig) error {
	if migrate == nil {
		return nil
	}
	repo, ok := repository.(*Repository)
	if !ok {
		return ErrorUnsupportedRepository
	}
	return migrate(repo.db, migrateEntities)
}

func (r *Repository) GetUsers(ids []int64) ([]*playground.User, error) {
	var users []*playground.User
	ctx := context.Background()
	err := r.db.NewSelect().Model(&users).Where("id = ?", In(ids)).Scan(ctx)
	return users, err
}

func (r *Repository) CreateUsers(entities []*playground.User) error {
	ctx := context.Background()
	users := make([]any, len(entities))
	for i, entity := range entities {
		f := NewUser
		_ = f
		users[i] = NewUser(entity)
	}
	return r.db.NewInsert().Model(&users).Scan(ctx)
}
