package app

import (
	"net/http"
	"playground/internal/entity"
)

var (
	app *application
)

func HealthCheck() int {
	return http.StatusOK
}

func Version() string {
	return "1.0.0-alpha"
}

var (
	authenticatePassword func(salt, password, correct string) (bool, error)
	encodePassword       func(salt []byte, password string) (string, error)
)

func GetUsers(ids []int64) ([]*entity.User, error) {
	return app.GetUsers(ids)
}

func CreateUsers(users []*entity.User) error {
	return app.CreateUsers(users)
}
