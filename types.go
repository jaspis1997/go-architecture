package playground

import "github.com/gin-gonic/gin"

type Repository interface {
	GetUsers(ids []int64) ([]*User, error)
	CreateUsers(users []*User) error
}

type RepositoryConfig interface{}

var (
	NewRepository func(RepositoryConfig) (Repository, error)
)

type (
	WebContext  = *gin.Context
	HandlerFunc = gin.HandlerFunc
)

var NewWebEngine func() Engine

type Engine interface {
	gin.IRouter
	Run(host string, port uint16, localhost ...bool) error
}
