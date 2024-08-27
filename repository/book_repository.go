package repository

import (
	"gorm.io/gorm"
	"library-management/models"
)

type BookRepository interface {
	Create(book *models.Book) error
	FindAll(queryParams map[string]string) ([]models.Book, error)
	FindById(id int) (*models.Book, error)
	Update(book *models.Book) error
	Delete(book *models.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (bookRepo *bookRepository) Create(book *models.Book) error {
	return bookRepo.db.Create(book).Error
}

func (bookRepo *bookRepository) FindAll(queryParams map[string]string) ([]models.Book, error) {
	var books []models.Book
	filterQuery := bookRepo.db

	// Filtering/Search by title / author
	if title, ok := queryParams["title"]; ok && title != "" {
		filterQuery = filterQuery.Where("title LIKE ?", "%"+title+"%")
	}
	if author, ok := queryParams["author"]; ok && author != "" {
		filterQuery = filterQuery.Where("author LIKE ?", "%"+author+"%")
	}

	// Filtering by price range
	if minPrice, ok := queryParams["minPrice"]; ok && minPrice != "" {
		filterQuery = filterQuery.Where("price >= ?", minPrice)
	}
	if maxPrice, ok := queryParams["maxPrice"]; ok && maxPrice != "" {
		filterQuery = filterQuery.Where("price <= ?", maxPrice)
	}

	// Sort by a specific field
	if sortBy, ok := queryParams["sortBy"]; ok && sortBy != "" {
		sortOrder := "asc"
		if sortDir, ok := queryParams["sortDir"]; ok && sortDir == "desc" {
			sortOrder = "desc"
		}
		filterQuery = filterQuery.Order(sortBy + " " + sortOrder)
	}

	err := filterQuery.Find(&books).Error
	return books, err
}

func (bookRepo *bookRepository) FindById(id int) (*models.Book, error) {
	var book models.Book
	err := bookRepo.db.Find(&book, id).Error
	return &book, err
}

func (bookRepo *bookRepository) Update(book *models.Book) error {
	return bookRepo.db.Save(book).Error
}

func (bookRepo *bookRepository) Delete(book *models.Book) error {
	return bookRepo.db.Delete(book).Error
}
