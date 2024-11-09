package repository

import (
	"context"
	"errors"
	playground "playground/internal"
	"playground/internal/entity"
	"reflect"
	"sync"
)

var (
	ErrorUnsupportedRepository = errors.New("unsupported repository")
)

type Repository struct {
	db Database
}

var (
	once    sync.Once
	models  map[string]any
	migrate func(db Database, targets []any) error
)

func migrateModels() []any {
	var models []any
	return models
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
