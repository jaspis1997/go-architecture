package main

import (
	"net/http"
	"playground"
	"playground/internal/web/handler"
)

func routes(e playground.Engine) playground.Engine {
	e.GET("/", func(c playground.WebContext) {
		c.Status(http.StatusOK)
	})
	e.GET(PathHealthCheck, handler.HealthCheck()...)
	e.GET(PathSystemVersion, handler.SystemVersion()...)

	e.GET("/api/v1/user/:id", handler.GetUsers()...)
	e.POST("/api/v1/user", handler.CreateUsers()...)

	e.GET("/favicon", func(c playground.WebContext) {
		c.File("../assets/favicon.png")
	})

	return e
}
