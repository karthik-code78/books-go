package handlers

import (
	"gorm.io/gorm"
	"library-management/middlewares/logging"
	"library-management/repository"
)

var db *gorm.DB

var AdminRepo repository.AdminRepository
var BooksRepo repository.BookRepository

func SetDatabase(database *gorm.DB) {
	logging.Log.Info(database)
	db = database

	AdminRepo = repository.NewAdminRepository(db)
	BooksRepo = repository.NewBookRepository(db)
}
