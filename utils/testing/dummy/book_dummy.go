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

func DummyRequestBook() domain.RequestBook {
	return domain.RequestBook{
		ISBN:        "dummy-isbn",
		Title:       "dummy-title",
		Author:      "dummy-author",
		Description: "dummy-description",
		PageCount:   100,
		CoverUrl:    "dummy-cover-url",
	}
}

func DummyResponseBook() domain.ResponseBook {
	return domain.ResponseBook{
		ISBN:        "dummy-isbn",
		Title:       "dummy-title",
		Author:      "dummy-author",
		Description: "dummy-description",
		PageCount:   100,
		CoverUrl:    "dummy-cover-url",
	}
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

func DummyResponseDetailBook() domain.ResponseDetailBook {
	return domain.ResponseDetailBook{
		ISBN:        "dummy-isbn",
		Title:       "dummy-title",
		Author:      "dummy-author",
		Description: "dummy-description",
		PageCount:   100,
		CoverUrl:    "dummy-cover-url",
		Reviews:     []domain.ResponseReview{},
	}
}
