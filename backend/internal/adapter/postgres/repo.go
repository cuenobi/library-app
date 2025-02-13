package postgres

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	return db
}
