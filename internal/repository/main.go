package repository

import (
	"errors"
	playground "playground/internal"

	"github.com/uptrace/bun"
)

var (
	ErrorUnsupportedConfig = errors.New("unsupported config")
)

type (
	Database = *bun.DB
)

var openDatabase func(DatabaseConfig) (Database, error)

func RegisterDatabaseDriver(f func(DatabaseConfig) (Database, error)) {
	openDatabase = f
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
