package handler

import (
	"net/http"
	"playground/internal/app"
	"playground/internal/web"
)

type (
	Context = web.Context
)

type DebugTokenStore int

func (DebugTokenStore) VerificationToken(token string) bool {
	return token == "debug"
}

func HealthCheck() []web.HandlerFunc {
	return []web.HandlerFunc{healthCheck}
}

func SystemVersion() []web.HandlerFunc {
	store := DebugTokenStore(0)
	return []web.HandlerFunc{
		web.AuthorizationBearerToken(store),
		web.RequestLimiter(),
		systemVersion,
	}
}

func healthCheck(c Context) {
	code := app.HealthCheck()
	c.Status(code)
}

func systemVersion(c Context) {
	version := app.Version()
	c.String(http.StatusOK, version)
}
