package mocks

import (
	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (m *BookRepositoryMock) Create(book *domain.Book) (*domain.Book, error) {
	args := m.Called(book)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) Delete(book *domain.Book) error {
	args := m.Called(book)

	return args.Error(0)
}

func (m *BookRepositoryMock) Fetch() ([]*domain.Book, error) {
	args := m.Called()

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) FetchBookReviews(book *domain.Book) ([]*domain.Review, error) {
	args := m.Called(book)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*domain.Review), args.Error(1)
}

func (m *BookRepositoryMock) FindByISBN(isbn string) (*domain.Book, error) {
	args := m.Called(isbn)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) FindByTitle(title string) ([]*domain.Book, error) {
	args := m.Called(title)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*domain.Book), args.Error(1)
}

func (m *BookRepositoryMock) Update(book *domain.Book) (*domain.Book, error) {
	args := m.Called(book)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*domain.Book), args.Error(1)
}
