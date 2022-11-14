package http

import (
	"net/http"

	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"github.com/fikriyusrihan/golang-clean-arch/interface/controller"
	"github.com/labstack/echo/v4"
)

func GetBooks(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.QueryParam("title")

		if title != "null" {
			return ctr.GetBooksByTitle(title, c)
		}

		return ctr.GetBooks(c)
	}
}

func GetBookByISBN(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		isbn := c.Param("isbn")

		return ctr.GetBookByISBN(isbn, c)
	}
}

func PostBook(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		bookRequest := new(domain.RequestBook)
		if err := c.Bind(&bookRequest); err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		return ctr.CreateBook(bookRequest, c)
	}
}

func PutBook(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		bookRequest := new(domain.RequestBook)
		if err := c.Bind(bookRequest); err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		return ctr.UpdateBook(bookRequest, c)
	}
}

func DeleteBook(ctr controller.AppController) echo.HandlerFunc {
	return func(c echo.Context) error {
		bookRequest := new(domain.RequestBook)
		if err := c.Bind(bookRequest); err != nil {
			return c.String(http.StatusBadRequest, "Bad Request")
		}

		book := domain.Book{
			ISBN: bookRequest.ISBN,
		}

		return ctr.DeleteBook(&book, c)
	}
}
