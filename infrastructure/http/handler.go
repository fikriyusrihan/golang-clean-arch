package http

import (
	"github.com/fikriyusrihan/golang-clean-arch/interface/controller"
	"github.com/labstack/echo/v4"
)

func GetBooks(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctr.GetBooks(c)
	}
}

func GetBookByISBN(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctr.GetBookByISBN(c)
	}
}

func PostBook(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctr.CreateBook(c)
	}
}

func PutBook(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctr.UpdateBook(c)
	}
}

func DeleteBook(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ctr.DeleteBook(c)
	}
}
