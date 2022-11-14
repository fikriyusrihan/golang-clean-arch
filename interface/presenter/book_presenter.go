package presenter

import "github.com/fikriyusrihan/golang-clean-arch/domain"

type bookPresenter struct{}

type BookPresenter interface {
	ResponseBooks(books []*domain.Book) []*domain.ResponseBook
	ResponseBook(book *domain.Book) *domain.ResponseDetailBook
	ResponseReviews(reviews []*domain.Review) []domain.ResponseReview
}

func NewBookPresenter() BookPresenter {
	return &bookPresenter{}
}

func (bp *bookPresenter) ResponseBooks(books []*domain.Book) []*domain.ResponseBook {
	var booksResponse []*domain.ResponseBook

	for _, book := range books {
		bookResponse := domain.ResponseBook{
			ISBN:        book.ISBN,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			PageCount:   book.PageCount,
			CoverUrl:    book.CoverUrl,
		}

		booksResponse = append(booksResponse, &bookResponse)
	}

	return booksResponse
}

func (bp *bookPresenter) ResponseBook(book *domain.Book) *domain.ResponseDetailBook {
	var reviews []domain.ResponseReview

	bookResponseDetail := domain.ResponseDetailBook{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		PageCount:   book.PageCount,
		CoverUrl:    book.CoverUrl,
		Reviews:     reviews,
	}

	return &bookResponseDetail
}

func (bp *bookPresenter) ResponseReviews(reviews []*domain.Review) []domain.ResponseReview {
	var reviewsResponse []domain.ResponseReview

	for _, review := range reviews {
		reviewResponse := domain.ResponseReview{
			ID:        review.ID,
			Writer:    review.Writer,
			Content:   review.Content,
			Rating:    review.Rating,
			CreatedAt: review.CreatedAt,
		}

		reviewsResponse = append(reviewsResponse, reviewResponse)
	}

	return reviewsResponse
}
