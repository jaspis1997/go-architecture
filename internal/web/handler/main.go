package handler

import (
	"net/http"
	"playground"
	"playground/internal/app"
	"playground/internal/web"
)

type (
	Context = playground.WebContext
)

type DebugTokenStore int

func (DebugTokenStore) IsActiveToken(token string) bool {
	return token == "debug"
}

func HealthCheck() []playground.HandlerFunc {
	return []playground.HandlerFunc{healthCheck}
}

func SystemVersion() []playground.HandlerFunc {
	store := DebugTokenStore(0)
	return []playground.HandlerFunc{
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
