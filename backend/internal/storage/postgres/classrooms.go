package postgres

import (
	"context"

  "github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *DB) CreateClassroom(ctx context.Context, classData models.Classroom) (error) {
  rows, err := db.connPool.Query(ctx,
    "INSERT INTO classrooms (ghclassroom_id, name, prof_id) VALUES ($1, $2, $3)", 
    classData.GHClassroom_ID, classData.Name, classData.Prof_ID)
  if err != nil {
    return err
  }


  defer rows.Close()
	return nil 
}

func (db *DB) GetAllClassrooms(ctx context.Context) ([]models.Classroom, error) {
  rows, err := db.connPool.Query(ctx, "SELECT ghclassroom_id, name, prof_id FROM classrooms")
  if err != nil {
    return nil, err 
  }

  defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Classroom])
}
