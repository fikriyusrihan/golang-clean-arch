package repository

import "github.com/fikriyusrihan/golang-clean-arch/domain"

type BookRepository interface {
	Create(book *domain.RequestBook) (*domain.Book, error)
	Update(book *domain.RequestBook) (*domain.Book, error)
	Delete(book *domain.RequestBook) error
	Fetch() ([]*domain.Book, error)
	FetchBookReviews(book *domain.Book) ([]*domain.Review, error)
	FindByISBN(isbn string) (*domain.Book, error)
	FindByTitle(title string) ([]*domain.Book, error)
}
