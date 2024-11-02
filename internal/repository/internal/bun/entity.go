package bun

import (
	"playground"
	"playground/internal/repository/internal/database"

	"github.com/uptrace/bun"
)

func MigrateEntities() []any {
	return []any{
		&User{},
	}
}

type User struct {
	bun.BaseModel `bun:"table:users"`
	Id            int64  `bun:",pk,autoincrement"`
	UniqueId      string `bun:"unique_id,notnull,unique,type:text"`
	Name          string `bun:"name,notnull"`
	Email         string `bun:"email,notnull,unique"`

	Salt     string `bun:"salt,notnull"`
	Password string `bun:"password"`
}

func NewUser(entity *playground.User) any {
	user := &User{}
	_ = database.ConvertDatabaseEntity(user, entity)
	return user
}

func (u *User) Entity() *playground.User {
	entity := &playground.User{}
	database.ConvertDatabaseEntity(u, entity)
	return entity
}
