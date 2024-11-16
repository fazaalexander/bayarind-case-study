package user

import "github.com/labstack/echo/v4"

func (uh *UserHandler) RegisterRoutes(e *echo.Echo) {
	authGroup := e.Group("/auth")
	authGroup.POST("/register", uh.Register())
	authGroup.POST("/login", uh.Login())
}
