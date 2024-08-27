package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"library-management/middlewares/logging"
	"library-management/utils"
	"os"
)

func ConnectAndReturnDB() (*gorm.DB, error) {
	logging.Log.Info("lol")
	utils.LoadEnv()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetJwtSecretKey() string {
	utils.LoadEnv()
	return os.Getenv("JWT_SECRET")
}
