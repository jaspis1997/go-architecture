package repository

type Config struct {
	Main DatabaseConfig
}

type DatabaseConfig interface {
	DSN() string
}
