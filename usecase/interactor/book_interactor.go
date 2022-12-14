package interactor

import (
	"time"

	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"github.com/fikriyusrihan/golang-clean-arch/usecase/presenter"
	"github.com/fikriyusrihan/golang-clean-arch/usecase/repository"
)

type bookInteractor struct {
	BookRepository repository.BookRepository
	BookPresenter  presenter.BookPresenter
}

type BookInteractor interface {
	Create(book *domain.RequestBook) (*domain.ResponseDetailBook, error)
	Update(book *domain.RequestBook) (*domain.ResponseDetailBook, error)
	Delete(book *domain.RequestBook) error
	Get() ([]*domain.ResponseBook, error)
	GetByISBN(isbn string) (*domain.ResponseDetailBook, error)
	GetByTitle(title string) ([]*domain.ResponseBook, error)
	GetBookReviews(book *domain.Book) ([]domain.ResponseReview, error)
}

func NewBookInteractor(r repository.BookRepository, p presenter.BookPresenter) BookInteractor {
	return &bookInteractor{r, p}
}

func (bi *bookInteractor) Create(bookRequest *domain.RequestBook) (*domain.ResponseDetailBook, error) {
	book := domain.Book{
		ISBN:        bookRequest.ISBN,
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
		PageCount:   bookRequest.PageCount,
		CoverUrl:    bookRequest.CoverUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	bookResult, err := bi.BookRepository.Create(&book)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBook(bookResult), nil
}

func (bi *bookInteractor) Update(bookRequest *domain.RequestBook) (*domain.ResponseDetailBook, error) {
	book := domain.Book{
		ISBN:        bookRequest.ISBN,
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
		PageCount:   bookRequest.PageCount,
		CoverUrl:    bookRequest.CoverUrl,
		UpdatedAt:   time.Now(),
	}

	bookResult, err := bi.BookRepository.Update(&book)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBook(bookResult), nil
}

func (bi *bookInteractor) Delete(bookRequest *domain.RequestBook) error {
	book := domain.Book{
		ISBN: bookRequest.ISBN,
	}

	return bi.BookRepository.Delete(&book)
}

func (bi *bookInteractor) Get() ([]*domain.ResponseBook, error) {
	books, err := bi.BookRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBooks(books), nil
}

func (bi *bookInteractor) GetByISBN(isbn string) (*domain.ResponseDetailBook, error) {
	book, err := bi.BookRepository.FindByISBN(isbn)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBook(book), nil
}

func (bi *bookInteractor) GetByTitle(title string) ([]*domain.ResponseBook, error) {
	books, err := bi.BookRepository.FindByTitle(title)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseBooks(books), nil
}

func (bi *bookInteractor) GetBookReviews(book *domain.Book) ([]domain.ResponseReview, error) {
	reviews, err := bi.BookRepository.FetchBookReviews(book)
	if err != nil {
		return nil, err
	}

	return bi.BookPresenter.ResponseReviews(reviews), nil
}
