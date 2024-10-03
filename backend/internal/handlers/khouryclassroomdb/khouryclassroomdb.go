package khouryclassroomdb

import (
  "net/http"
	"github.com/gofiber/fiber/v2"
)

func (s *KCDBService) GetUsers(c *fiber.Ctx) error {
  users, err := s.store.GetUsers(c.Context())
  if err != nil {
    return err
  }

  return c.Status(http.StatusOK).JSON(users)
  
}


func (s *KCDBService) CreateClassroom(c *fiber.Ctx) error {

  var classData Classroom
  err := c.BodyParser(&classData)
  if err != nil {
    return InvalidJSON() 
  }


  err := s.store.CreateClassroom(c.Context(), classData)
}
