package app

import (
	"errors"
	"net/http"
	"playground"
	"playground/internal/app/crypto"
	"sync"
)

var (
	ErrorUnsupportedRepository = errors.New("unsupported repository")
	ErrorEmptyUsers            = errors.New("empty users")
)

func HealthCheck() int {
	return http.StatusOK
}

func Version() string {
	return "1.0.0-alpha"
}

type app struct {
	*users
}

var (
	authenticatePassword func(salt, password, correct string) (bool, error)
	encodePassword       func(salt []byte, password string) (string, error)
)

var (
	initLock    sync.Mutex
	application playground.App
)

func newApp(repo playground.Repository) playground.App {
	return &app{
		users: &users{Repo: repo},
	}
}

func Init(repo playground.Repository) error {
	initLock.Lock()
	defer initLock.Unlock()
	if repo == nil {
		return ErrorUnsupportedRepository
	}
	if application == nil {
		application = newApp(repo)
		err := initModules()
		if err != nil {
			return err
		}
	}
	return nil
}

func initModules() error {
	options := crypto.NewDefaultOptions()
	authenticatePassword = crypto.AuthenticatePassword(options)
	encodePassword = crypto.EncodePassword(options)
	return nil
}

type users struct {
	Repo playground.Repository
}

func (u *users) GetUsers(ids []int64) ([]*playground.User, error) {
	return u.Repo.GetUsers(ids)
}

func (u *users) CreateUsers(users []*playground.User) error {
	if len(users) == 0 {
		return ErrorEmptyUsers
	}
	salt, err := crypto.GenerateRandomSalt(crypto.DefaultSaltLength)
	if err != nil {
		return err
	}
	for _, user := range users {
		user.UniqueId = crypto.GenerateUUID()
		user.Salt = crypto.EncodeSalt(salt)
		user.Password, err = encodePassword(salt, user.Password)
		if err != nil {
			return err
		}
	}
	return u.Repo.CreateUsers(users)
}

func GetUsers(ids []int64) ([]*playground.User, error) {
	return application.GetUsers(ids)
}

func CreateUsers(users []*playground.User) error {
	return application.CreateUsers(users)
}
