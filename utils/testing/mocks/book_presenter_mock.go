package mocks

import (
	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type BookPresenterMock struct {
	mock.Mock
}

func (m *BookPresenterMock) ResponseBook(book *domain.Book) *domain.ResponseDetailBook {
	args := m.Called(book)

	return args.Get(0).(*domain.ResponseDetailBook)
}

func (m *BookPresenterMock) ResponseBooks(books []*domain.Book) []*domain.ResponseBook {
	args := m.Called(books)

	return args.Get(0).([]*domain.ResponseBook)
}

func (m *BookPresenterMock) ResponseReviews(reviews []*domain.Review) []domain.ResponseReview {
	args := m.Called(reviews)

	return args.Get(0).([]domain.ResponseReview)
}
