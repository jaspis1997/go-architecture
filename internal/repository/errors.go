package repository

import (
	"errors"
)

var (
	ErrorUnsupportedConfig     = errors.New("unsupported config")
	ErrorUnsupportedRepository = errors.New("unsupported repository")
)
