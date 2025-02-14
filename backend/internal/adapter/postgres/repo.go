package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"library-service/internal/constant"
	"library-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

// NewPostgres creates a new connection to the PostgreSQL database.
// It takes a PostgresConfig struct pointer, a context, and returns a pointer to the connection.
// The connection is deferred to be closed after the function returns.
func NewPostgres(cfg *PostgresConfig, ctx context.Context) *gorm.DB {
	// Construct the PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	// Open the connection
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("connect to db error:", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.BorrowDetail{}, &model.Book{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	return db
}

// SeedData populates the database with some initial data.
// The function takes a pointer to the connection as an argument.
func SeedData(db *gorm.DB) {
	// Create two users: John Doe (admin) and Jane Doe (user).
	users := []model.User{
		{
			ID:        uuid.New().String(),
			Username:  "john_doe",
			Password:  "password123",
			Name:      "John Doe",
			Role:      "admin",
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
		{
			ID:        uuid.New().String(),
			Username:  "jane_doe",
			Password:  "password123",
			Name:      "Jane Doe",
			Role:      "user",
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
	}

	// Insert the users into the database. If the username already exists,
	// updates the existing user with the given data.
	for _, user := range users {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}}, // The conflict should be checked based on the username.
			DoUpdates: clause.AssignmentColumns([]string{"username", "password", "name", "role", "created_at"}), // The columns that should be updated if a conflict occurs.
		}).Create(&user).Error; err != nil {
			panic("failed to seed users: " + err.Error())
		}
	}

	// Create two books: "The Go Programming Language" and "Clean Code".
	books := []model.Book{
		{
			Name:      "The Go Programming Language",
			Category:  "Programming",
			Status:    constant.Available,
			Stock:     3,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
		{
			Name:      "Clean Code",
			Category:  "Programming",
			Status:    constant.Available,
			Stock:     5,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
	}

	// Insert the books into the database. If the book name already exists,
	// updates the existing book with the given data.
	for _, book := range books {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}}, // The conflict should be checked based on the book name.
			DoUpdates: clause.AssignmentColumns([]string{"name", "category", "created_at"}), // The columns that should be updated if a conflict occurs.
		}).Create(&book).Error; err != nil {
			panic("failed to seed books: " + err.Error())
		}
	}
}
