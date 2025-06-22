package handlers

import (
	"net/http"

	"github.com/justchatapp/justchat/services"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers HTTP routes for the server.
func RegisterRoutes(e *echo.Echo) {
	e.GET("/", index)
	e.POST("/register", register)
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to JustChat!")
}

func register(c echo.Context) error {
	var req struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := services.UserService.Register(c.Request().Context(), req.Username, req.Password); err != nil {
		return err
	}
	return c.String(http.StatusOK, "registered")
}
