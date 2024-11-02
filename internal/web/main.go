package web

import (
	"fmt"
	"playground"

	"github.com/gin-gonic/gin"
)

type (
	HandlerFunc = playground.HandlerFunc
	Context     = playground.WebContext
)

type Engine struct {
	*gin.Engine
}

func (e *Engine) Run(host string, port uint16, localhost ...bool) error {
	addr := func() string {
		if len(localhost) > 0 && localhost[0] {
			return fmt.Sprintf("localhost:%d", port)
		}
		return fmt.Sprintf("%s:%d", host, port)
	}
	gin.SetMode(gin.ReleaseMode)
	return e.Engine.Run(addr())
}

func New() playground.Engine {
	return &Engine{gin.New()}
}
