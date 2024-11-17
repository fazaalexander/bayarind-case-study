package author

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (ah *AuthorHandler) RegisterRoutes(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	authorGroup := e.Group("/authors")
	authorGroup.Use(jwtMiddleware)
	authorGroup.GET("", ah.GetAllAuthor())
	authorGroup.GET("/:id", ah.GetAuthorById())
	authorGroup.POST("", ah.CreateAuthor())
	authorGroup.PUT("/:id", ah.UpdateAuthor())
	authorGroup.DELETE("/:id", ah.DeleteAuthor())
}
