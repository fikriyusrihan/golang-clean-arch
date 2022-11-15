package http

import (
	"github.com/fikriyusrihan/golang-clean-arch/interface/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, ctr controller.AppController) *echo.Echo {
	logFormat := "[${time_custom}] ${method} ${host}${uri} [${status} ${latency_human}] ${error}\n"
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Format: logFormat, CustomTimeFormat: "2006-01-02 15:04:05.00000"}))
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")
	{
		v1.GET("/books", GetBooks(ctr))
		v1.GET("/books/:isbn", GetBookByISBN(ctr))
		v1.POST("/books", PostBook(ctr))
		v1.PUT("/books/:isbn", PutBook(ctr))
		v1.DELETE("/books/:isbn", DeleteBook(ctr))
	}

	return e
}
