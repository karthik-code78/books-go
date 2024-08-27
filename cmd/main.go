package main

import (
	"library-management/config"
	"library-management/database"
	"library-management/handlers"
	"library-management/middlewares/logging"
	"library-management/router"
	"log"
	"net/http"
)

func main() {
	logging.Initializelogger()
	//var log = logging.LoggerInst
	db, err := config.ConnectAndReturnDB()
	if err != nil {
		logging.Log.Fatal("Failed to connect to the database", err)
	}

	logging.Log.Info("db is", db)

	// Initialize tables
	database.Migrate()

	// Database for handlers
	handlers.SetDatabase(db)

	// Init router
	mainRouter := router.InitRouter()

	// Start server
	//log.Info("Server is running on port : 8080")
	err = http.ListenAndServe(":8080", mainRouter)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
