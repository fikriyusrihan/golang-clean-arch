package http

import (
	"github.com/fikriyusrihan/golang-clean-arch/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, ctr controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", GetBooks(ctr))
	e.GET("/books/:isbn", GetBookByISBN(ctr))
	e.POST("/books", PostBook(ctr))
	e.PUT("/books/:isbn", PutBook(ctr))
	e.DELETE("/books/:isbn", DeleteBook(ctr))

	return e
}
