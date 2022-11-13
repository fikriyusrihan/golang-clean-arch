package presenter

import "github.com/fikriyusrihan/golang-clean-arch/domain"

type BookPresenter interface {
	ResponseBooks(books []*domain.Book) []*domain.ResponseBook
	ResponseBook(book *domain.Book) *domain.ResponseDetailBook
}