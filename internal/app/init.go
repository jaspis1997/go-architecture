package app

import (
	"playground"
	"playground/internal/app/crypto"
	"sync"
)

var once sync.Once

func Init(repo playground.Repository) error {
	if repo == nil {
		return ErrorUnsupportedRepository
	}
	once.Do(initialize(repo))
	return nil
}

func initialize(repo playground.Repository) func() {
	return func() {
		if app != nil {
			panic(ErrorInitializedApplication)
		}
		err := initApplication(repo)
		if err != nil {
			panic(err)
		}
		err = initModules()
		if err != nil {
			panic(err)
		}
	}
}

func initApplication(repo playground.Repository) error {
	if repo == nil {
		return ErrorUnsupportedRepository
	}
	app = &application{
		users: &users{Repo: repo},
	}
	return nil
}

func initModules() error {
	options := crypto.NewDefaultOptions()
	authenticatePassword = crypto.AuthenticatePassword(options)
	encodePassword = crypto.EncodePassword(options)
	return nil
}
