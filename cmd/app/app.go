package app

import (
	"github.com/fazaalexander/bayarind-case-study.git/cmd/routes"
	"github.com/fazaalexander/bayarind-case-study.git/common"
	"github.com/fazaalexander/bayarind-case-study.git/database/mysql"

	userHandler "github.com/fazaalexander/bayarind-case-study.git/modules/handler/api/user"
	userRepo "github.com/fazaalexander/bayarind-case-study.git/modules/repository/user"
	userUsecase "github.com/fazaalexander/bayarind-case-study.git/modules/usecase/user"

	authorHandler "github.com/fazaalexander/bayarind-case-study.git/modules/handler/api/author"
	authorRepo "github.com/fazaalexander/bayarind-case-study.git/modules/repository/author"
	authorUsecase "github.com/fazaalexander/bayarind-case-study.git/modules/usecase/author"

	bookHandler "github.com/fazaalexander/bayarind-case-study.git/modules/handler/api/book"
	bookRepo "github.com/fazaalexander/bayarind-case-study.git/modules/repository/book"
	bookUsecase "github.com/fazaalexander/bayarind-case-study.git/modules/usecase/book"

	"github.com/labstack/echo/v4"
)

func StartApp() *echo.Echo {
	mysql.Init()

	userRepo := userRepo.New(mysql.DB)
	userUsecase := userUsecase.New(userRepo)
	userHandler := userHandler.New(userUsecase)

	authorRepo := authorRepo.New(mysql.DB)
	authorUsecase := authorUsecase.New(authorRepo)
	authorHandler := authorHandler.New(authorUsecase)

	bookRepo := bookRepo.New(mysql.DB)
	bookUsecase := bookUsecase.New(bookRepo)
	bookHandler := bookHandler.New(bookUsecase)

	handler := common.Handler{
		UserHandler:   userHandler,
		AuthorHandler: authorHandler,
		BookHandler:   bookHandler,
	}

	router := routes.StartRoute(handler)

	return router
}
