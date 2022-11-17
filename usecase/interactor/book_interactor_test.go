package interactor_test

import (
	"errors"
	"testing"

	"github.com/fikriyusrihan/golang-clean-arch/usecase/interactor"
	"github.com/fikriyusrihan/golang-clean-arch/utils/testing/dummy"
	"github.com/fikriyusrihan/golang-clean-arch/utils/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	dummyBook := dummy.DummyBook()
	dummyBookRequest := dummy.DummyRequestBook()
	dummyBookResponse := dummy.DummyResponseDetailBook()

	t.Run("Create - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Create", mock.Anything).Return(&dummyBook, nil).Once()
		mockBookPresenter.On("ResponseBook", &dummyBook).Return(&dummyBookResponse).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		bookResponse, err := uc.Create(&dummyBookRequest)

		assert.NoError(t, err)
		assert.NotNil(t, bookResponse)
		assert.Equal(t, dummyBookRequest.ISBN, bookResponse.ISBN)
		assert.Equal(t, dummyBookRequest.Title, bookResponse.Title)

		mockBookRepository.AssertCalled(t, "Create", mock.Anything)
		mockBookPresenter.AssertCalled(t, "ResponseBook", &dummyBook)

		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("Create - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Create", mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		bookResponse, err := uc.Create(&dummyBookRequest)

		assert.Error(t, err)
		assert.Nil(t, bookResponse)

		mockBookRepository.AssertCalled(t, "Create", mock.Anything)
		mockBookRepository.AssertNotCalled(t, "ResponseBook", mock.AnythingOfType("*entity.Book"))
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	dummyBook := dummy.DummyBook()
	dummyBookRequest := dummy.DummyRequestBook()
	dummyBookResponse := dummy.DummyResponseDetailBook()

	t.Run("Update - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Update", mock.Anything).Return(&dummyBook, nil).Once()
		mockBookPresenter.On("ResponseBook", &dummyBook).Return(&dummyBookResponse).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.Update(&dummyBookRequest)

		assert.NoError(t, err)
		assert.NotNil(t, responseBook)
		assert.Equal(t, dummyBookRequest.Author, responseBook.Author)

		mockBookRepository.AssertCalled(t, "Update", mock.Anything)
		mockBookPresenter.AssertCalled(t, "ResponseBook", &dummyBook)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("Update - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Update", mock.Anything).Return(nil, errors.New("Unexpected error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.Update(&dummyBookRequest)

		assert.Error(t, err)
		assert.Nil(t, responseBook)

		mockBookRepository.AssertCalled(t, "Update", mock.Anything)
		mockBookRepository.AssertNotCalled(t, "ResponseBook", mock.AnythingOfType("*entity.Book"))
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	dummyBookRequest := dummy.DummyRequestBook()

	t.Run("Delete - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Delete", mock.Anything).Return(nil).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		err := uc.Delete(&dummyBookRequest)

		assert.NoError(t, err)

		mockBookRepository.AssertCalled(t, "Delete", mock.Anything)
		mockBookRepository.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	dummyBooks := dummy.DummyBooks()
	dummyResponseBooks := dummy.DummyResponseBooks()

	t.Run("Get - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("Fetch").Return(dummyBooks, nil).Once()
		mockBookPresenter.On("ResponseBooks", dummyBooks).Return(dummyResponseBooks)

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBooks, err := uc.Get()

		assert.NoError(t, err)
		assert.NotNil(t, responseBooks)
		assert.Equal(t, dummyBooks[0].Title, responseBooks[0].Title)

		mockBookRepository.AssertCalled(t, "Fetch")
		mockBookPresenter.AssertCalled(t, "ResponseBooks", dummyBooks)
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
	dummyBook := dummy.DummyBook()
	dummyBookResponse := dummy.DummyResponseDetailBook()

	t.Run("GetByISBN - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByISBN", dummyBook.ISBN).Return(&dummyBook, nil).Once()
		mockBookPresenter.On("ResponseBook", &dummyBook).Return(&dummyBookResponse).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByISBN(dummyBook.ISBN)

		assert.NoError(t, err)
		assert.NotNil(t, responseBook)
		assert.Equal(t, dummyBook.ISBN, responseBook.ISBN)

		mockBookRepository.AssertCalled(t, "FindByISBN", dummyBook.ISBN)
		mockBookPresenter.AssertCalled(t, "ResponseBook", &dummyBook)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("GetByISBN - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByISBN", dummyBook.ISBN).Return(nil, errors.New("Unexpected Error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByISBN(dummyBook.ISBN)

		assert.Nil(t, responseBook)
		assert.Error(t, err)

		mockBookRepository.AssertCalled(t, "FindByISBN", dummyBook.ISBN)
		mockBookPresenter.AssertNotCalled(t, "ResponseBook", &dummyBook)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestGetByTitle(t *testing.T) {
	dummyBook := dummy.DummyBook()
	dummyBooks := dummy.DummyBooks()
	dummyResponseBooks := dummy.DummyResponseBooks()

	t.Run("GetByTitle - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByTitle", dummyBook.Title).Return(dummyBooks, nil).Once()
		mockBookPresenter.On("ResponseBooks", dummyBooks).Return(dummyResponseBooks).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByTitle(dummyBook.Title)

		assert.NoError(t, err)
		assert.NotNil(t, responseBook)

		mockBookRepository.AssertCalled(t, "FindByTitle", dummyBook.Title)
		mockBookPresenter.AssertCalled(t, "ResponseBooks", dummyBooks)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("GetByTitle - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FindByTitle", dummyBook.Title).Return(nil, errors.New("Unexpected Error")).Once()

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseBook, err := uc.GetByTitle(dummyBook.Title)

		assert.Nil(t, responseBook)
		assert.Error(t, err)

		mockBookRepository.AssertCalled(t, "FindByTitle", dummyBook.Title)
		mockBookPresenter.AssertNotCalled(t, "ResponseBooks", dummyBooks)
		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}

func TestGetReviews(t *testing.T) {
	dummyBook := dummy.DummyBook()
	dummyBookReviews := dummy.DummyBookReviews()
	dummyResponseReviews := dummy.DummyBookResponseReviews()

	t.Run("GetReviews - Success", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FetchBookReviews", &dummyBook).Return(dummyBookReviews, nil)
		mockBookPresenter.On("ResponseReviews", dummyBookReviews).Return(dummyResponseReviews, nil)

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseReview, err := uc.GetBookReviews(&dummyBook)

		assert.NoError(t, err)
		assert.NotNil(t, responseReview)

		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})

	t.Run("GetReviews - Failed", func(t *testing.T) {
		mockBookRepository := new(mocks.BookRepositoryMock)
		mockBookPresenter := new(mocks.BookPresenterMock)

		mockBookRepository.On("FetchBookReviews", &dummyBook).Return(nil, errors.New("Unexpected Error"))

		uc := interactor.NewBookInteractor(mockBookRepository, mockBookPresenter)

		responseReview, err := uc.GetBookReviews(&dummyBook)

		assert.Nil(t, responseReview)
		assert.Error(t, err)

		mockBookRepository.AssertExpectations(t)
		mockBookPresenter.AssertExpectations(t)
	})
}
