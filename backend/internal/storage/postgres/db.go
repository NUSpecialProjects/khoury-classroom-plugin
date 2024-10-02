package postgres

import (
	"context"
	"fmt"
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}


// Establishes a database connection and returns it for querying.
func New(ctx context.Context, config config.Database) (*DB, error) {
	conn, err := pgx.Connect(ctx, config.URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	
  fmt.Println("Successfully connected to the database!")

  return &DB{conn: conn}, nil
}


func (db *DB) Close(ctx context.Context) error {
	return db.conn.Close(ctx)
}


func (db *DB) GetUsers(ctx contect.Context) error {
  rows, err := db.conn.Query(ctx, "SELECT * FROM users")
  if err != nil {
    return []models.User{}, err
  }

  defer rows.Close()
  return pgx


}
