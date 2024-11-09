package crypto

import "errors"

var (
	ErrorUnsupported     = errors.New("unsupported encryption algorithm")
	ErrorInvalidSalt     = errors.New("invalid salt")
	ErrorInvalidPassword = errors.New("invalid password")

	ErrorInvalidSaltLength = errors.New("invalid salt length")
)
