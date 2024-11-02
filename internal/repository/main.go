package repository

import (
	"errors"
	"playground"

	"github.com/uptrace/bun"
)

var (
	ErrorUnsupportedConfig = errors.New("unsupported config")
)

type (
	Database = *bun.DB
)

var newDatabase func(DatabaseConfig) (Database, error)

func Register(f func(DatabaseConfig) (Database, error)) {
	newDatabase = f
}

func New(config playground.RepositoryConfig) (playground.Repository, error) {
	conf, ok := config.(Config)
	if !ok {
		return nil, ErrorUnsupportedConfig
	}
	db, err := newDatabase(conf.Main)
	if err != nil {
		return nil, err
	}
	repo := &Repository{
		db: db,
	}
	return repo, nil
}
