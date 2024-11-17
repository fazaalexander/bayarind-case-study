package routes

import (
	"github.com/fazaalexander/bayarind-case-study.git/common"
	"github.com/fazaalexander/bayarind-case-study.git/middleware/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartRoute(handler common.Handler) *echo.Echo {
	e := echo.New()
	log.LogMiddleware(e)
	e.Use(middleware.CORS())

	handler.UserHandler.RegisterRoutes(e)
	handler.AuthorHandler.RegisterRoutes(e)
	handler.BookHandler.RegisterRoutes(e)

	return e
}
