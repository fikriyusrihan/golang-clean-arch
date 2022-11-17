package interactor_test

import (
	"errors"
	"testing"

	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"github.com/fikriyusrihan/golang-clean-arch/usecase/interactor"
	"github.com/fikriyusrihan/golang-clean-arch/utils/testing/dummy"
	"github.com/fikriyusrihan/golang-clean-arch/utils/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockBook := dummy.DummyBook()
	mockBookRequest := dummy.DummyRequestBook()
	mockBookResponse := dummy.DummyResponseDetailBook()

	t.Run("Create - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Create", mock.Anything).Return(&mockBook, nil).Once()
		mockBookPresenter.On("ResponseBook", &mockBook).Return(&mockBookResponse).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		bookResponse, err := uc.Create(&mockBookRequest)

		assert.NoError(t, err)
		assert.NotNil(t, bookResponse)
		assert.Equal(t, mockBookRequest.ISBN, bookResponse.ISBN)
		assert.Equal(t, mockBookRequest.Title, bookResponse.Title)

		mockBookRepository.AssertCalled(t, "Create", mock.Anything)
		mockBookPresenter.AssertCalled(t, "ResponseBook", &mockBook)

		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("Create - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Create", mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		bookResponse, err := uc.Create(&mockBookRequest)

		assert.Error(t, err)
		assert.Nil(t, bookResponse)

		mockBookRepository.AssertCalled(t, "Create", mock.Anything)
		mockBookRepository.AssertNotCalled(t, "ResponseBook", mock.AnythingOfType("*entity.Book"))
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockBook := dummy.DummyBook()
	mockBookRequest := dummy.DummyRequestBook()
	mockBookResponse := dummy.DummyResponseDetailBook()

	t.Run("Update - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Update", mock.Anything).Return(&mockBook, nil).Once()
		mockBookPresenter.On("ResponseBook", &mockBook).Return(&mockBookResponse).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.Update(&mockBookRequest)

		assert.NoError(t, err)
		assert.NotNil(t, responseBook)
		assert.Equal(t, mockBookRequest.Author, responseBook.Author)

		mockBookRepository.AssertCalled(t, "Update", mock.Anything)
		mockBookPresenter.AssertCalled(t, "ResponseBook", &mockBook)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("Update - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Update", mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.Update(&mockBookRequest)

		assert.Error(t, err)
		assert.Nil(t, responseBook)

		mockBookRepository.AssertCalled(t, "Update", mock.Anything)
		mockBookRepository.AssertNotCalled(t, "ResponseBook", mock.AnythingOfType("*entity.Book"))
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockBookRequest := dummy.DummyRequestBook()

	t.Run("Delete - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Delete", mock.Anything).Return(nil).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		err := uc.Delete(&mockBookRequest)

		assert.NoError(t, err)

		mockBookRepository.AssertCalled(t, "Delete", mock.Anything)
		mockBookRepository.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	mockBook := dummy.DummyBook()
	mockResponseBook := dummy.DummyResponseBook()

	mockBooks := []*domain.Book{
		&mockBook,
	}

	mockResponseBooks := []*domain.ResponseBook{
		&mockResponseBook,
	}

	t.Run("Get - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Fetch").Return(mockBooks, nil).Once()
		mockBookPresenter.On("ResponseBooks", mockBooks).Return(mockResponseBooks)

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBooks, err := uc.Get()

		assert.NoError(t, err)
		assert.NotNil(t, responseBooks)
		assert.Equal(t, mockBooks[0].Title, responseBooks[0].Title)

		mockBookRepository.AssertCalled(t, "Fetch")
		mockBookPresenter.AssertCalled(t, "ResponseBooks", mockBooks)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("Get - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Fetch").Return(nil, errors.New("Unexpected error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBooks, err := uc.Get()

		assert.Error(t, err)
		assert.Nil(t, responseBooks)

		mockBookRepository.AssertCalled(t, "Fetch")
		mockBookRepository.AssertNotCalled(t, "ResponseBooks", mock.AnythingOfType("[]*entity.Book"))
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestGetByISBN(t *testing.T) {
	mockBook := dummy.DummyBook()
	mockBookResponse := dummy.DummyResponseDetailBook()

	t.Run("GetByISBN - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByISBN", mockBook.ISBN).Return(&mockBook, nil).Once()
		mockBookPresenter.On("ResponseBook", &mockBook).Return(&mockBookResponse).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByISBN(mockBook.ISBN)

		assert.NoError(t, err)
		assert.NotNil(t, responseBook)
		assert.Equal(t, mockBook.ISBN, responseBook.ISBN)

		mockBookRepository.AssertCalled(t, "FindByISBN", mockBook.ISBN)
		mockBookPresenter.AssertCalled(t, "ResponseBook", &mockBook)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("GetByISBN - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByISBN", mockBook.ISBN).Return(nil, errors.New("Unexpected Error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByISBN(mockBook.ISBN)

		assert.Nil(t, responseBook)
		assert.Error(t, err)

		mockBookRepository.AssertCalled(t, "FindByISBN", mockBook.ISBN)
		mockBookPresenter.AssertNotCalled(t, "ResponseBook", &mockBook)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestGetByTitle(t *testing.T) {
	mockBook := dummy.DummyBook()
	mockResponseBook := dummy.DummyResponseBook()

	mockBooks := []*domain.Book{
		&mockBook,
	}

	mockResponseBooks := []*domain.ResponseBook{
		&mockResponseBook,
	}

	t.Run("GetByTitle - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByTitle", mockBook.Title).Return(mockBooks, nil).Once()
		mockBookPresenter.On("ResponseBooks", mockBooks).Return(mockResponseBooks).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByTitle(mockBook.Title)

		assert.NoError(t, err)
		assert.NotNil(t, responseBook)

		mockBookRepository.AssertCalled(t, "FindByTitle", mockBook.Title)
		mockBookPresenter.AssertCalled(t, "ResponseBooks", mockBooks)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("GetByTitle - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByTitle", mockBook.Title).Return(nil, errors.New("Unexpected Error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByTitle(mockBook.Title)

		assert.Nil(t, responseBook)
		assert.Error(t, err)

		mockBookRepository.AssertCalled(t, "FindByTitle", mockBook.Title)
		mockBookPresenter.AssertNotCalled(t, "ResponseBooks", mockBooks)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestGetReviews(t *testing.T) {
	mockBook := dummy.DummyBook()
	mockBookReview := dummy.DummyBookReview()
	mockBookResponseReview := dummy.DummyBookResponseReview()

	mockBookReviews := []*domain.Review{
		&mockBookReview,
	}

	mockResponseReviews := []domain.ResponseReview{
		mockBookResponseReview,
	}

	t.Run("GetReviews - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FetchBookReviews", &mockBook).Return(mockBookReviews, nil)
		mockBookPresenter.On("ResponseReviews", mockBookReviews).Return(mockResponseReviews, nil)

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseReview, err := uc.GetBookReviews(&mockBook)

		assert.NoError(t, err)
		assert.NotNil(t, responseReview)

		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("GetReviews - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FetchBookReviews", &mockBook).Return(nil, errors.New("Unexpected Error"))

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseReview, err := uc.GetBookReviews(&mockBook)

		assert.Nil(t, responseReview)
		assert.Error(t, err)

		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}
