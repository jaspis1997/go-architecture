package web

import (
	"net/http"
	"strings"

	playground "playground/internal"

	"golang.org/x/time/rate"
)

func AuthorizationBearerToken(store playground.TokenAuthorization) HandlerFunc {
	return func(c Context) {
		token := c.GetHeader(HeaderAuthorization)
		token = strings.TrimPrefix(token, PrefixAuthorizationToken)
		if !store.VerificationToken(token) {
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
