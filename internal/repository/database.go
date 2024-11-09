package repository

import (
	playground "playground/internal"
	"reflect"
	"sync"

	"github.com/uptrace/bun"
)

type (
	Database = *bun.DB
)

var (
	once    sync.Once
	models  map[string]any
	migrate func(db Database, targets []any) error
)

var openDatabase func(DatabaseConfig) (Database, error)

func RegisterDatabaseDriver(f func(DatabaseConfig) (Database, error)) {
	openDatabase = f
}

func RegisterModels(model []any) {
	once.Do(func() {
		models = make(map[string]any)
	})
	for _, v := range model {
		models[reflect.TypeOf(v).Elem().Name()] = v
	}
}

func ConvertModel(entity any) any {
	model := models[reflect.TypeOf(entity).Elem().Name()]
	ConvertDatabaseEntity(model, entity)
	return model
}

func RegisterMigrate(f func(db Database, targets []any) error) {
	migrate = f
}

func migrateModels() []any {
	var m []any
	for _, v := range models {
		m = append(m, v)
	}
	return m
}

func Migrate(repository playground.RepositoryConfig) error {
	if migrate == nil {
		// TODO: define error migrate is nil
		panic("TODO")
	}
	repo, ok := repository.(*Repository)
	if !ok {
		return ErrorUnsupportedRepository
	}
	return migrate(repo.db, migrateModels())
}
