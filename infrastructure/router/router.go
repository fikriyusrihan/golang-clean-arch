package router

import (
	"github.com/fikriyusrihan/golang-clean-arch/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/books", func(context echo.Context) error { return c.GetBooks(context) })
	e.GET("/books/:isbn", func(context echo.Context) error {
		isbn := context.Param("isbn")
		
		return c.GetBookByISBN(isbn, context)
	})

	return e
}
