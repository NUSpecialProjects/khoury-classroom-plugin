package khouryclassroomdb

import (
  "net/http"
	"github.com/gofiber/fiber/v2"
  "github.com/CamPlume1/khoury-classroom/internal/models"
  "github.com/CamPlume1/khoury-classroom/internal/errs"
)


func (s *KCDBService) GetUsers(c *fiber.Ctx) error {
  users, err := s.store.GetUsers(c.Context())
  if err != nil {
    return err
  }

  return c.Status(http.StatusOK).JSON(users)
  
}

func (s *KCDBService) GetAllClassrooms(c *fiber.Ctx) error {
  classrooms, err := s.store.GetAllClassrooms(c.Context())
  if err != nil {
    return err
  }

  return c.Status(http.StatusOK).JSON(classrooms)
  
}

func (s *KCDBService) CreateClassroom(c *fiber.Ctx) error {

  var classData models.Classroom
  err := c.BodyParser(&classData)
  if err != nil {
    return errs.InvalidJSON() 
  }


  error := s.store.CreateClassroom(c.Context(), classData)
  if error != nil {
    return error
  }

  c.Status(http.StatusOK)

  return nil
}
