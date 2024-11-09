package handler

import (
	"net/http"
	playground "playground/internal"
	"playground/internal/app"
	"playground/internal/entity"
	"playground/internal/web"
	"strconv"
)

func GetUsers() []web.HandlerFunc {
	store := DebugTokenStore(0)
	return []web.HandlerFunc{
		web.AuthorizationBearerToken(store),
		web.RequestLimiter(),
		getUsers,
	}
}

func getUsers(c Context) {
	id, err := strconv.Atoi(c.Param(playground.KeyId))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	users, err := app.GetUsers([]int64{int64(id)})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(users) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, users[0])
}

func CreateUsers() []web.HandlerFunc {
	store := DebugTokenStore(0)
	return []web.HandlerFunc{
		web.AuthorizationBearerToken(store),
		web.RequestLimiter(),
		createUsers,
	}
}

func createUsers(c Context) {
	var user entity.User
	c.Bind(&user)
	err := app.CreateUsers([]*entity.User{&user})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
