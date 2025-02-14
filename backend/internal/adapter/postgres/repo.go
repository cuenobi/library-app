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

func SeedData(db *gorm.DB) {
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

	for _, user := range users {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}}, //
			DoUpdates: clause.AssignmentColumns([]string{"username", "password", "name", "role", "created_at"}),
		}).Create(&user).Error; err != nil {
			panic("failed to seed users: " + err.Error())
		}
	}

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

	for _, book := range books {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "category", "created_at"}),
		}).Create(&book).Error; err != nil {
			panic("failed to seed books: " + err.Error())
		}
	}
}
