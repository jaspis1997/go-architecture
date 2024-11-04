package app

import "errors"

var (
	ErrorUnsupportedRepository  = errors.New("unsupported repository")
	ErrorEmptyUsers             = errors.New("empty users")
	ErrorInitializedApplication = errors.New("application already initialized")
)
