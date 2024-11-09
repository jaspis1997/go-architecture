package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	HandlerFunc = gin.HandlerFunc
	Context     = *gin.Context
)

type Engine interface {
	gin.IRouter
	Run(host string, port uint16) error
}

type engine struct {
	*gin.Engine
}

func (e *engine) Run(host string, port uint16) error {
	addr := func() string {
		return fmt.Sprintf("%s:%d", host, port)
	}
	return e.Engine.Run(addr())
}

func New() Engine {
	return &engine{gin.New()}
}
