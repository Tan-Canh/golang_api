package routes

import (
	"github.com/labstack/echo/v4"
	"golang_api/handlers"
	"net/http"
)

type API struct {
	Echo *echo.Echo
	UserHandler handlers.HandlerUser
}

func (api *API) SetupRouter() {
	//ping
	api.Echo.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ping OK")
	})

	//user
	api.Echo.GET("/users", api.UserHandler.GetList)
	api.Echo.POST("/users", api.UserHandler.Add)
	api.Echo.GET("/users/:id", api.UserHandler.GetUserById)
	api.Echo.PUT("/users/:id", api.UserHandler.Update)
}
