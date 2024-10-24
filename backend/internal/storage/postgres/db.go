package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	connPool *pgxpool.Pool
}

// Establishes a postgres connection pool and returns it for querying
func New(ctx context.Context, config config.Database) (*DB, error) {
	connPool, err := pgxpool.New(ctx, config.URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return &DB{connPool: connPool}, nil
}

// Closes the connection pool
func (db *DB) Close(ctx context.Context) {
	db.connPool.Close()
}

// Loads and executes a SQL file
func (db *DB) ExecFile(ctx context.Context, filePath string) error {
	// Read the migration file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read migration file '%s': %w", filePath, err)
	}

	// Execute the migration SQL
	_, err = db.connPool.Exec(ctx, string(content))
	if err != nil {
		return fmt.Errorf("failed to execute migration '%s': %w", filePath, err)
	}

	log.Printf("Successfully applied migration: %s", filePath)
	return nil
}
