package app

import (
	playground "playground/internal"
	"playground/internal/app/crypto"
	"playground/internal/entity"
)

type application struct {
	*users
}

type users struct {
	Repo playground.Repository
}

func (u *users) GetUsers(ids []int64) ([]*entity.User, error) {
	return u.Repo.GetUsers(ids)
}

func (u *users) CreateUsers(users []*entity.User) error {
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
