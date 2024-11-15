package common

import (
	ah "github.com/fazaalexander/bayarind-case-study.git/modules/handler/api/author"
	bh "github.com/fazaalexander/bayarind-case-study.git/modules/handler/api/book"
	uh "github.com/fazaalexander/bayarind-case-study.git/modules/handler/api/user"
)

type Handler struct {
	UserHandler   *uh.UserHandler
	AuthorHandler *ah.AuthorHandler
	BookHandler   *bh.BookHandler
}
