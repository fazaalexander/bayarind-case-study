package book

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (bh *BookHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	bookGroup := e.Group("/books")
	bookGroup.Use(jwtMiddleware)
	bookGroup.GET("", bh.GetAllBooks())
	bookGroup.GET("/:id", bh.GetBookById())
	bookGroup.POST("", bh.CreateBook())
	bookGroup.PUT("/:id", bh.UpdateBook())
	bookGroup.DELETE("/:id", bh.DeleteBook())
}
