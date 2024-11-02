package init

import (
	"playground/internal/repository"
	"playground/internal/repository/internal/bun"
)

func init() {
	repository.Register(bun.New)
	repository.RegisterMigrate(bun.Migrate)
	repository.RegisterEntities(bun.MigrateEntities())

	repository.NewUser = bun.NewUser
}
