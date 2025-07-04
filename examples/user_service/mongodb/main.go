package main

import (
	"errors"
	"github.com/Lumicrate/gompose/core"
	"github.com/Lumicrate/gompose/db/mongodb"
	"github.com/Lumicrate/gompose/http/gin"
	"github.com/Lumicrate/gompose/http/middlewares"
	"log"
	"strings"
	"time"
)

// User entity definition with basic fields
type User struct {
	ID    int    `json:"id" bson:"id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

// Define the hooks for validation or any other puposes
func (u *User) BeforeCreate() error {
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

func (u *User) AfterDelete() error {
	log.Printf("User %d deleted", u.ID)
	return nil
}

func main() {

	// MongoDB URI and database
	mongoURI := "mongodb://localhost:27017"
	dbName := "users"

	// Initialize Mongo adapter
	dbAdapter := mongodb.New(mongoURI, dbName)

	// Initialize Gin HTTP adapter on port 8080
	httpEngine := ginadapter.New(8080)

	// Create app instance
	app := core.NewApp().
		AddEntity(User{}).
		UseDB(dbAdapter).
		RegisterMiddleware(middlewares.LoggingMiddleware()).
		RegisterMiddleware(middlewares.RateLimitMiddleware(1 * time.Second)).
		UseHTTP(httpEngine)

	// Run the app
	log.Println("Starting service on :8080")
	app.Run()
}
