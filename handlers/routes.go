package handlers

import (
	"net/http"

	"github.com/justchatapp/justchat/services"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes registers HTTP routes for the server.
func RegisterRoutes(e *echo.Echo) {
	e.GET("/", loginPage)
	e.GET("/login", loginPage)
	e.POST("/login", login)
	e.GET("/register", registerPage)
	e.POST("/register", register)
	e.GET("/chat", chatPage)
}

func loginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func registerPage(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

func chatPage(c echo.Context) error {
	return c.Render(http.StatusOK, "chat.html", nil)
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
	return c.Redirect(http.StatusSeeOther, "/")
}

func login(c echo.Context) error {
	var req struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}

	ok, err := services.UserService.Authenticate(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return err
	}
	if !ok {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}
	return c.Redirect(http.StatusSeeOther, "/chat")
}
