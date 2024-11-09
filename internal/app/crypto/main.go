package crypto

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

const (
	DefaultMemory     = 64 * 1024
	DefaultTime       = 1
	DefaultThreads    = 4
	DefaultKeyLength  = 32
	DefaultSaltLength = 32
)

type Options interface {
}

type Argon2Options struct {
	Memory    uint32
	Time      uint32
	Threads   uint8
	KeyLength uint32
}

func NewDefaultOptions() Options {
	return Argon2Options{
		Memory:    DefaultMemory,
		Time:      DefaultTime,
		Threads:   DefaultThreads,
		KeyLength: DefaultKeyLength,
	}
}

func AuthenticatePassword(opt Options) func(salt, password, correct string) (bool, error) {
	return func(salt, password, correct string) (bool, error) {
		return VerifyPassword(salt, password, correct, opt)
	}
}

func EncodePassword(opt Options) func(salt []byte, password string) (string, error) {
	return func(salt []byte, password string) (string, error) {
		return HashPassword(salt, password, opt)
	}
}

func VerifyPassword(salt, password, correct string, opt Options) (bool, error) {
	hash, err := func() (string, error) {
		salt, err := DecodeSalt(salt)
		if err != nil {
			return "", err
		}
		return HashPassword(salt, password, opt)
	}()
	if err != nil {
		return false, err
	}
	return hash == correct, nil
}

func HashPassword(salt []byte, password string, opt Options) (string, error) {
	switch opt := opt.(type) {
	case Argon2Options:
		return hashPasswordArgon2(password, salt, opt)
	default:
		return "", ErrorUnsupported
	}
}

func GenerateRandomSalt(length int) ([]byte, error) {
	if length <= 0 {
		return nil, ErrorInvalidSaltLength
	}
	salt := make([]byte, length)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func EncodeSalt(salt []byte) string {
	return base64.StdEncoding.EncodeToString(salt)
}

func DecodeSalt(salt string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(salt)
}

func validateSalt(salt []byte) error {
	if len(salt) != DefaultSaltLength {
		return ErrorInvalidSalt
	}
	return nil
}

func hashPasswordArgon2(password string, salt []byte, opt Argon2Options) (string, error) {
	if err := validateSalt(salt); err != nil {
		return "", err
	}
	var (
		memory    = opt.Memory
		time      = opt.Time
		threads   = opt.Threads
		keyLength = opt.KeyLength
	)
	result := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)
	return base64.StdEncoding.EncodeToString(result), nil
}

func GenerateUUID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return uuid.String()
}
