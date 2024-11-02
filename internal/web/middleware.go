package web

import (
	"net/http"
	"strings"

	"golang.org/x/time/rate"
)

type TokenStore interface {
	IsActiveToken(token string) bool
}

func AuthorizationBearerToken(store TokenStore) HandlerFunc {
	return func(c Context) {
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		if !store.IsActiveToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func RequestLimiter() HandlerFunc {
	limiter := rate.NewLimiter(2, 1)
	return func(c Context) {
		if limiter.Allow() {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusTooManyRequests)
	}
}
