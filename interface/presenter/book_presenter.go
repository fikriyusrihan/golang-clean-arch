package presenter

import "github.com/fikriyusrihan/golang-clean-arch/domain"

type bookPresenter struct{}

type BookPresenter interface {
	ResponseBooks(books []*domain.Book) []*domain.ResponseBook
	ResponseBook(book *domain.Book) *domain.ResponseDetailBook
}

func NewBookPresenter() BookPresenter {
	return &bookPresenter{}
}

func (bp *bookPresenter) ResponseBooks(books []*domain.Book) []*domain.ResponseBook {
	var booksResponse []*domain.ResponseBook

	for _, book := range books {
		bookResponse := domain.ResponseBook{
			ISBN: book.ISBN,
			Title: book.Title,
			Author: book.Author,
			Description: book.Description,
			PageCount: book.PageCount,
			CoverUrl: book.CoverUrl,
		}

		booksResponse = append(booksResponse, &bookResponse)
	}

	return booksResponse
}

func (bp *bookPresenter) ResponseBook(book *domain.Book) *domain.ResponseDetailBook {
	var reviews []domain.Review

	bookResponseDetail := domain.ResponseDetailBook {
		ISBN: book.ISBN,
		Title: book.Title,
		Author: book.Author,
		Description: book.Description,
		PageCount: book.PageCount,
		CoverUrl: book.CoverUrl,
		Reviews: reviews,
	}

	return &bookResponseDetail
}
