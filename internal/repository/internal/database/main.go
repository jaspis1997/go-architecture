package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"playground/internal/repository"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type (
	Config   = repository.DatabaseConfig
	Database = repository.Database

	PostgresConfig = repository.PostgresConfig
	SQLiteConfig   = repository.SQLiteConfig
)

func New(config Config) (Database, error) {
	return open(config)
}

const (
	FormatErrorRequired = "required field is empty: %s"
)

var (
	ErrorUnsupportedConfig = errors.New("unsupported config")
)

func open(config Config) (Database, error) {
	switch config := config.(type) {
	case nil:
		return open(PostgresConfig{})
	case PostgresConfig:
		dsn := config.DSN()
		sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		db := bun.NewDB(sqldb, pgdialect.New())
		return db, nil
	case *PostgresConfig:
		return open(*config)
	case SQLiteConfig:
		if config.Filename == "" {
			return nil, fmt.Errorf(FormatErrorRequired, "config.Filename")
		}
		sqldb, err := sql.Open(sqliteshim.ShimName, config.Filename)
		if err != nil {
			return nil, err
		}
		return bun.NewDB(sqldb, sqlitedialect.New()), nil
	case *SQLiteConfig:
		return open(*config)
	}
	return nil, ErrorUnsupportedConfig
}

func Migrate(db Database, targets []any) error {
	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, target := range targets {
		_, err := tx.NewCreateTable().
			Model(target).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}
