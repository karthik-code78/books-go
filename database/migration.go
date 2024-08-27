package database

import (
	"library-management/config"
	"library-management/middlewares/logging"
	"library-management/models"
)

func Migrate() {
	db, err := config.ConnectAndReturnDB()
	if err != nil {
		logging.Log.Fatal("Failed to connect to the Database", err)
	}
	logging.Log.Info("DB in migrate is: ", db)
	err = db.AutoMigrate(&models.Book{}, &models.Admin{})
	if err != nil {
		logging.Log.Error("Failed to migrate DB", err)
	}
}
