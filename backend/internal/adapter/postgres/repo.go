package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
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
func NewPostgres(cfg *PostgresConfig, ctx context.Context) *pgx.Conn {
	// Construct the PostgreSQL connection string
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	// Open the connection
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	// Test the connection
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal("Ping failed:", err)
	}

	// Defer the closing of the connection
	defer conn.Close(ctx)

	return conn
}
