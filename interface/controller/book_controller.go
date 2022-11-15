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
		apiResponse := domain.ResponseApiError{
			Error:   true,
			Message: "Bad Request",
		}

		return c.JSON(http.StatusBadRequest, apiResponse)
	}

	bookDetail, err := bc.bookInteractor.Create(bookRequest)
	if err != nil {
		return err
	}

	apiResponse := domain.ResponseApiSuccess{
		Error:   false,
		Message: "success",
		Data:    bookDetail,
	}

	return c.JSON(http.StatusCreated, apiResponse)
}

func (bc *bookController) GetBooks(c Context) error {
	title := c.QueryParam("title")
	if title != "null" {
		return bc.GetBooksByTitle(c)
	}

	books, err := bc.bookInteractor.Get()
	if err != nil {
		return err
	}

	apiResponse := domain.ResponseApiSuccess{
		Error:   false,
		Message: "success",
		Data:    books,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (bc *bookController) DeleteBook(c Context) error {
	bookRequest := new(domain.RequestBook)
	if err := c.Bind(bookRequest); err != nil {
		apiResponse := domain.ResponseApiError{
			Error:   true,
			Message: "Bad Request",
		}

		return c.JSON(http.StatusBadRequest, apiResponse)
	}

	if err := bc.bookInteractor.Delete(bookRequest); err != nil {
		return err
	}

	apiResponse := domain.ResponseApiSuccess{
		Error:   false,
		Message: "success",
		Data: map[string]string{
			"isbn": bookRequest.ISBN,
		},
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (bc *bookController) GetBookByISBN(c Context) error {
	isbn := c.Param("isbn")

	book, err := bc.bookInteractor.GetByISBN(isbn)
	if err != nil {
		return err
	}

	reviews, err := bc.bookInteractor.GetBookReviews(&domain.Book{
		ISBN: book.ISBN,
	})
	if err != nil {
		return err
	}

	book.Reviews = reviews

	apiResponse := domain.ResponseApiSuccess{
		Error:   false,
		Message: "success",
		Data:    book,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (bc *bookController) GetBooksByTitle(c Context) error {
	title := c.Param("title")
	
	book, err := bc.bookInteractor.GetByTitle(title)
	if err != nil {
		return err
	}

	apiResponse := domain.ResponseApiSuccess{
		Error:   false,
		Message: "success",
		Data:    book,
	}

	return c.JSON(http.StatusOK, apiResponse)
}

func (bc *bookController) UpdateBook(c Context) error {
	bookRequest := new(domain.RequestBook)
	if err := c.Bind(bookRequest); err != nil {
		apiResponse := domain.ResponseApiError{
			Error:   true,
			Message: "Bad Request",
		}

		return c.JSON(http.StatusBadRequest, apiResponse)
	}

	bookResponse, err := bc.bookInteractor.Update(bookRequest)
	if err != nil {
		return err
	}

	apiResponse := domain.ResponseApiSuccess{
		Error:   false,
		Message: "success",
		Data:    bookResponse,
	}

	return c.JSON(http.StatusOK, apiResponse)
}
