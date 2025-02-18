package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"library-service/internal/constant"
	"library-service/internal/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostgresConfig struct {
	Host         string
	Port         string
	Name         string
	Username     string
	Password     string
	SeedPassword string
}

// NewPostgres creates a new connection to the PostgreSQL database.
// It takes a PostgresConfig struct pointer, a context, and returns a pointer to the connection.
// The connection is deferred to be closed after the function returns.
func NewPostgres(cfg *PostgresConfig, ctx context.Context) *gorm.DB {
	// Construct the PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)

	var db *gorm.DB
	var err error

	// Retry up to 5 times
	for i := 0; i < 5; i++ {
		// Open the connection
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err == nil {
			// Successfully connected, proceed with migration
			break
		}

		// If connection fails, log the error and wait before retrying
		log.Printf("Failed to connect to DB (attempt %d/5): %v", i+1, err)
		time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
	}

	if err != nil {
		log.Fatal("Could not connect to DB after 5 attempts:", err)
	}

	// Perform database migrations
	if err := db.AutoMigrate(&model.User{}, &model.BorrowDetail{}, &model.Book{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	return db
}

// SeedData populates the database with some initial data.
// The function takes a pointer to the connection as an argument.
func SeedData(cfg *PostgresConfig, db *gorm.DB) {
	// Hash password
	password := cfg.SeedPassword
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Create two users: John Doe (admin) and Jane Doe (user).
	users := []model.User{
		{
			ID:        uuid.New().String(),
			Username:  "john_doe",
			Password:  string(hashedPassword),
			Name:      "John Doe",
			Role:      constant.Librarian,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
		{
			ID:        uuid.New().String(),
			Username:  "jane_doe",
			Password:  string(hashedPassword),
			Name:      "Jane Doe",
			Role:      constant.MemberRole,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
	}

	// Insert the users into the database.
	for _, user := range users {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}},
			DoUpdates: clause.AssignmentColumns([]string{"username", "password", "name", "role", "created_at"}),
		}).Create(&user).Error; err != nil {
			panic("failed to seed users: " + err.Error())
		}
	}

	// Create books
	books := []model.Book{
		{
			ID:        uuid.New().String(),
			Name:      "The Go Programming Language",
			Category:  "Programming",
			Status:    constant.Available,
			Stock:     3,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
		{
			ID:        uuid.New().String(),
			Name:      "Clean Code",
			Category:  "Programming",
			Status:    constant.Available,
			Stock:     5,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		},
	}

	// Insert books
	for _, book := range books {
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "category", "created_at"}),
		}).Create(&book).Error; err != nil {
			panic("failed to seed books: " + err.Error())
		}
	}

	var count int64
	db.Model(&model.BorrowDetail{}).Count(&count)
	if count > 0 {
		log.Println("Borrow details already exist, skipping seed data")
		return // ถ้ามีข้อมูลอยู่แล้ว ให้ออกจากฟังก์ชัน ไม่ต้องเพิ่มข้อมูลซ้ำ
	}

	// หา Jane Doe
	var janeDoe model.User
	if err := db.Where("username = ?", "jane_doe").First(&janeDoe).Error; err != nil {
		panic("failed to find jane_doe user: " + err.Error())
	}

	// หา Books
	var booksExist []model.Book
	if err := db.Find(&booksExist).Error; err != nil {
		panic("failed to find books: " + err.Error())
	}

	// BorrowDetails
	var borrowDetails []*model.BorrowDetail
	for _, book := range booksExist {
		borrowDetails = append(borrowDetails, &model.BorrowDetail{
			BookName:   book.Name,
			BorrowedAt: time.Now().Unix(),
			UserID:     janeDoe.ID,
			BookID:     book.ID,
		})
	}

	// Insert BorrowDetails
	for _, borrow := range borrowDetails {
		if err := db.Create(&borrow).Error; err != nil {
			panic("failed to seed borrow details: " + err.Error())
		}
	}

	log.Println("Borrow details seeded successfully")
}
