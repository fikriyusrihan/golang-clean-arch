package dummy

import (
	"time"

	"github.com/fikriyusrihan/golang-clean-arch/domain"
)

func DummyBook() domain.Book {
	return domain.Book{
		ISBN:        "dummy-isbn",
		Title:       "dummy-title",
		Author:      "dummy-author",
		Description: "dummy-description",
		PageCount:   100,
		CoverUrl:    "dummy-cover-url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func DummyBooks() []*domain.Book {
	book := DummyBook()
	return []*domain.Book{&book}
}

func DummyRequestBook() domain.RequestBook {
	book := DummyBook()
	return domain.RequestBook{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		PageCount:   book.PageCount,
		CoverUrl:    book.CoverUrl,
	}
}

func DummyResponseBook() domain.ResponseBook {
	book := DummyBook()
	return domain.ResponseBook{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		PageCount:   book.PageCount,
		CoverUrl:    book.CoverUrl,
	}
}

func DummyResponseBooks() []*domain.ResponseBook {
	book := DummyResponseBook()
	return []*domain.ResponseBook{&book}
}

func DummyBookReview() domain.Review {
	return domain.Review{
		ID:        1,
		ISBN:      "dummy-isbn",
		Writer:    "dummy-writer",
		Content:   "dummy-content",
		Rating:    5,
		CreatedAt: time.Now(),
	}
}

func DummyBookReviews() []*domain.Review {
	review := DummyBookReview()
	return []*domain.Review{&review}
}

func DummyBookResponseReview() domain.ResponseReview {
	review := DummyBookReview()
	return domain.ResponseReview{
		ID:        review.ID,
		Writer:    review.Writer,
		Content:   review.Content,
		Rating:    review.Rating,
		CreatedAt: review.CreatedAt,
	}
}

func DummyBookResponseReviews() []domain.ResponseReview {
	review := DummyBookResponseReview()
	return []domain.ResponseReview{review}
}

func DummyResponseDetailBook() domain.ResponseDetailBook {
	book := DummyBook()
	return domain.ResponseDetailBook{
		ISBN:        book.ISBN,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		PageCount:   book.PageCount,
		CoverUrl:    book.CoverUrl,
		Reviews:     []domain.ResponseReview{},
	}
}
