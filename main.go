package main

import (
	"blankproject/entities"
	"blankproject/httpLocal"
	"blankproject/repository"
	"blankproject/usecases"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const urlServer = ":8080"

func main() {
	db, err := gorm.Open(sqlite.Open("blank.db"), &gorm.Config{})
	ErrorHandling(err, db)

	// Define the message creation
	messageRepo := &repository.GormMessageRepository{DB: db}
	messageUseCase := &usecases.RepoMessageUseCase{REPO: messageRepo}
	messageHandler := &httpLocal.MessageHandler{SERVICE: messageUseCase}

	// Define the user creation
	userRepo := &repository.GormUserRepository{DB: db}
	userUseCase := &usecases.RepoUserUseCase{REPO: userRepo}
	userHandler := &httpLocal.UserHandler{SERVICE: userUseCase}

	// Define the HTTP server and routes.
	http.HandleFunc("/message", messageHandler.CreateMessage)
	http.HandleFunc("/messages", messageHandler.GetAllMessages)
	http.HandleFunc("/user", userHandler.CreateUser)

	// Start the server.
	log.Printf("Server running on %s", urlServer)
	if err := http.ListenAndServe(urlServer, nil); err != nil {
		log.Fatalf("Could not start server: %s", err)
	}
}

// HTTP and Database connection Error hangling fuction
// - Error
// - DB
func ErrorHandling(err error, db *gorm.DB) {
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	if err := db.AutoMigrate(&entities.Message{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %s", err)
	}

	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %s", err)
	}
}
