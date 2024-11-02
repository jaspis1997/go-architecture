package init

import (
	"playground"
	"playground/internal/repository"
	"playground/internal/web"
)

func init() {
	playground.NewWebEngine = web.New
	playground.NewRepository = repository.New
}
