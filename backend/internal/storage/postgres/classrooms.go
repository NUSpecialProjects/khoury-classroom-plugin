package postgres

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
)


func (db *DB) CreateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error) {
    err := db.connPool.QueryRow(ctx, "INSERT INTO classrooms (name, org_id, org_name, created_at) VALUES ($1, $2, $3, $4) RETURNING *",
        classroomData.Name, 
        classroomData.OrgID,
        classroomData.OrgName,
        classroomData.CreatedAt,
        ).Scan(&classroomData.ID,
        &classroomData.Name,
        &classroomData.OrgID,
        &classroomData.OrgName,
        &classroomData.CreatedAt)

    if err != nil {
        return models.Classroom{}, errs.NewDBError(err)
    }
    
    return classroomData, nil
}


func (db *DB) UpdateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error) {

    err := db.connPool.QueryRow(ctx, "UPDATE classrooms SET name = $1, org_id = $2, org_name = $3 WHERE id = $4 RETURNING *",
        classroomData.Name, 
        classroomData.OrgID,
        classroomData.OrgName,
        classroomData.ID).Scan(&classroomData.ID,
        &classroomData.Name,
        &classroomData.OrgID,
        &classroomData.OrgName,
        &classroomData.CreatedAt)

    if err != nil {
        return models.Classroom{}, errs.NewDBError(err)
    }
    
    return classroomData, nil
}


func (db *DB) GetClassroomByID(ctx context.Context, classroomID int64) (models.Classroom, error) {
    var classroomData models.Classroom
    err := db.connPool.QueryRow(ctx, "SELECT * FROM classrooms WHERE id = $1", classroomID).Scan(
        &classroomData.ID,
        &classroomData.Name,
        &classroomData.OrgID,
        &classroomData.OrgName,
        &classroomData.CreatedAt,
        )

    if err != nil {
        return models.Classroom{}, errs.NewDBError(err)
    }

    return classroomData, nil
}

