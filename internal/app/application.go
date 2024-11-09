package app

import (
	playground "playground/internal"
	"playground/internal/app/crypto"
	"playground/internal/entity"
	"sync"
)

var waitGroup sync.WaitGroup

// Done waits for all database operations to finish.
func Done() {
	waitGroup.Wait()
}

type application struct {
	*users
}

type users struct {
	Repo playground.Repository
}

func (u *users) GetUsers(ids []int64) ([]*entity.User, error) {
	waitGroup.Add(1)
	defer waitGroup.Done()
	return u.Repo.GetUsers(ids)
}

func (u *users) CreateUsers(users []*entity.User) error {
	if len(users) == 0 {
		return ErrorEmptyUsers
	}
	waitGroup.Add(1)
	defer waitGroup.Done()
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
