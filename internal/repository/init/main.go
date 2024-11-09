package init

import (
	"playground/internal/repository"
	"playground/internal/repository/internal/bun"
)

func init() {
	repository.RegisterDatabaseDriver(bun.New)
	repository.RegisterMigrate(bun.Migrate)
	repository.RegisterModels(bun.MigrateEntities())

	// repository.NewUser = bun.NewUser
}
