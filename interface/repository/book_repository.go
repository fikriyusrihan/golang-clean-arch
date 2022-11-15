package repository

import (
	"time"

	"github.com/fikriyusrihan/golang-clean-arch/domain"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

type BookRepository interface {
	Create(book *domain.RequestBook) (*domain.Book, error)
	Update(book *domain.RequestBook) (*domain.Book, error)
	Delete(book *domain.RequestBook) error
	Fetch() ([]*domain.Book, error)
	FindByISBN(isbn string) (*domain.Book, error)
	FindByTitle(title string) ([]*domain.Book, error)
	FetchBookReviews(book *domain.Book) ([]*domain.Review, error)
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) Create(bookRequest *domain.RequestBook) (*domain.Book, error) {
	book := domain.Book{
		ISBN:        bookRequest.ISBN,
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
		PageCount:   bookRequest.PageCount,
		CoverUrl:    bookRequest.CoverUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := br.db.Create(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (br *bookRepository) Update(bookRequest *domain.RequestBook) (*domain.Book, error) {
	book := domain.Book{
		ISBN:        bookRequest.ISBN,
		Title:       bookRequest.Title,
		Author:      bookRequest.Author,
		Description: bookRequest.Description,
		PageCount:   bookRequest.PageCount,
		CoverUrl:    bookRequest.CoverUrl,
	}

	result := br.db.Save(book)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (br *bookRepository) Delete(bookRequest *domain.RequestBook) error {
	book := domain.Book{
		ISBN: bookRequest.ISBN,
	}

	result := br.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (br *bookRepository) Fetch() ([]*domain.Book, error) {
	var books []*domain.Book

	result := br.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (br *bookRepository) FetchBookReviews(book *domain.Book) ([]*domain.Review, error) {
	var reviews []*domain.Review

	result := br.db.Find(&reviews, "isbn = ?", book.ISBN)

	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}

func (br *bookRepository) FindByISBN(isbn string) (*domain.Book, error) {
	var book domain.Book

	result := br.db.First(&book, "isbn = ?", isbn)

	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (br *bookRepository) FindByTitle(title string) ([]*domain.Book, error) {
	var books []*domain.Book

	result := br.db.Where("title LIKE ?", "%"+title+"%").Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}
