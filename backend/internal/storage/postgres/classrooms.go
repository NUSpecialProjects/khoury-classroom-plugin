package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)


func (db *DB) AddClassroom(ctx context.Context, classData models.Classroom) error {
  err := db.connPool.Query(ctx,
    "INSERT INTO classrooms (ghclassroom_id, name, prof_id) VALUES ($1, $2, $3) RETURNING id", 
    classData.ghclassroom_id, classData.name, classData.prof_id)

  if err != nil {
    return nil
  }

}
