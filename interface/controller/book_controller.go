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
	CreateBook(book *domain.Book, c Context) error
	UpdateBook(book *domain.Book, c Context) error
	DeleteBook(book *domain.Book, c Context) error
	GetBooks(c Context) error
	GetBookByISBN(isbn string, c Context) error
	GetBooksByTitle(title string, c Context) error
}

func NewBookController(bi interactor.BookInteractor) BookController {
	return &bookController{bi}
}

func (bc *bookController) CreateBook(book *domain.Book, c Context) error {
	bookDetail, err := bc.bookInteractor.Create(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, bookDetail)
}

func (bc *bookController) GetBooks(c Context) error {
	books, err := bc.bookInteractor.Get()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, books)
}

func (bc *bookController) DeleteBook(book *domain.Book, c Context) error {
	if err := bc.bookInteractor.Delete(book); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) GetBookByISBN(isbn string, c Context) error {
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

	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) GetBooksByTitle(title string, c Context) error {
	book, err := bc.bookInteractor.GetByTitle(title)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (bc *bookController) UpdateBook(book *domain.Book, c Context) error {
	bookResponse, err := bc.bookInteractor.Update(book)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, bookResponse)
}
