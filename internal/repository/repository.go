package repository

import (
	"context"
	playground "playground/internal"
	"playground/internal/entity"
)

type Repository struct {
	db Database
}

func New(config playground.RepositoryConfig) (playground.Repository, error) {
	conf, ok := config.(Config)
	if !ok {
		return nil, ErrorUnsupportedConfig
	}
	db, err := openDatabase(conf.Main)
	if err != nil {
		return nil, err
	}
	repo := &Repository{
		db: db,
	}
	return repo, nil
}

func (r *Repository) GetUsers(ids []int64) ([]*entity.User, error) {
	var users []*entity.User
	ctx := context.Background()
	err := r.db.NewSelect().Model(&users).Where("id = ?", In(ids)).Scan(ctx)
	return users, err
}

func (r *Repository) CreateUsers(entities []*entity.User) error {
	ctx := context.Background()
	users := make([]any, len(entities))
	for i, entity := range entities {
		users[i] = ConvertModel(entity)
	}
	return r.db.NewInsert().Model(&users).Scan(ctx)
}
