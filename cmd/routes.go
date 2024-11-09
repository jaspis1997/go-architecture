package main

import (
	"net/http"
	"playground/internal/web"
	"playground/internal/web/handler"
)

func routes(e web.Engine) web.Engine {
	e.GET("/", func(c web.Context) {
		c.Status(http.StatusOK)
	})
	e.GET(PathHealthCheck, handler.HealthCheck()...)
	e.GET(PathSystemVersion, handler.SystemVersion()...)

	e.GET("/api/v1/user/:id", handler.GetUsers()...)
	e.POST("/api/v1/user", handler.CreateUsers()...)

	e.GET("/favicon", func(c web.Context) {
		c.File("../assets/favicon.png")
	})

	return e
}
