package controller

import (
	"net/http"

	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"github.com/fikriyusrihan/golang-clean-arch/usecase/interactor"
)

type bookController struct {
	bookInteractor interactor.BookInteractor
}

type BookController interface {
	CreateBook(c Context) error
	UpdateBook(c Context) error
	DeleteBook(c Context) error
	GetBooks(c Context) error
	GetBookByISBN(c Context) error
	GetBooksByTitle(c Context) error
}

func NewBookController(bi interactor.BookInteractor) BookController {
	return &bookController{bi}
}

func (bc *bookController) CreateBook(c Context) error {
	bookRequest := new(domain.RequestBook)
	if err := c.Bind(&bookRequest); err != nil {
		return errorApiResponse(http.StatusBadRequest, "Bad Request", c)
	}

	bookDetail, err := bc.bookInteractor.Create(bookRequest)
	if err != nil {
		return errorApiResponse(http.StatusInternalServerError, "Internal Server Error", c)
	}

	return successApiResponse(http.StatusCreated, bookDetail, c)
}

func (bc *bookController) GetBooks(c Context) error {
	title := c.QueryParam("title")
	if title != "null" {
		return bc.GetBooksByTitle(c)
	}

	books, err := bc.bookInteractor.Get()
	if err != nil {
		return errorApiResponse(http.StatusInternalServerError, "Internal Server Error", c)
	}

	return successApiResponse(http.StatusOK, books, c)
}

func (bc *bookController) DeleteBook(c Context) error {
	bookRequest := new(domain.RequestBook)
	if err := c.Bind(bookRequest); err != nil {
		return errorApiResponse(http.StatusBadRequest, "Bad Request", c)
	}

	if err := bc.bookInteractor.Delete(bookRequest); err != nil {
		return errorApiResponse(http.StatusInternalServerError, "Internal Server Error", c)
	}

	data := map[string]string{
		"isbn": bookRequest.ISBN,
	}

	return successApiResponse(http.StatusOK, data, c)
}

func (bc *bookController) GetBookByISBN(c Context) error {
	isbn := c.Param("isbn")

	book, err := bc.bookInteractor.GetByISBN(isbn)
	if err != nil {
		return errorApiResponse(http.StatusInternalServerError, "Internal Server Error", c)
	}

	reviews, err := bc.bookInteractor.GetBookReviews(&domain.Book{
		ISBN: book.ISBN,
	})
	if err != nil {
		return errorApiResponse(http.StatusInternalServerError, "Internal Server Error", c)
	}

	book.Reviews = reviews

	return successApiResponse(http.StatusOK, book, c)
}

func (bc *bookController) GetBooksByTitle(c Context) error {
	title := c.Param("title")

	book, err := bc.bookInteractor.GetByTitle(title)
	if err != nil {
		return errorApiResponse(http.StatusInternalServerError, "Internal Server Error", c)
	}

	return successApiResponse(http.StatusOK, book, c)
}

func (bc *bookController) UpdateBook(c Context) error {
	bookRequest := new(domain.RequestBook)
	if err := c.Bind(bookRequest); err != nil {
		return errorApiResponse(http.StatusBadRequest, "Bad Request", c)
	}

	bookResponse, err := bc.bookInteractor.Update(bookRequest)
	if err != nil {
		return errorApiResponse(http.StatusInternalServerError, "Internal Server Error", c)
	}

	return successApiResponse(http.StatusOK, bookResponse, c)
}

func errorApiResponse(code int, message string, c Context) error {
	apiResponse := domain.ResponseApiError{
		Error:   true,
		Message: message,
	}

	return c.JSON(code, apiResponse)
}

func successApiResponse(code int, data interface{}, c Context) error {
	apiResponse := domain.ResponseApiSuccess{
		Error:   false,
		Message: "Success",
		Data:    data,
	}

	return c.JSON(code, apiResponse)
}
