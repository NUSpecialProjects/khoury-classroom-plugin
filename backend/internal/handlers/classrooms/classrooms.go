package classrooms

import (
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (s *ClassroomService) GetUsersInClassroom(c *fiber.Ctx) error {
  classroomID := c.Params("classroomID")

	users, err := s.store.GetUsersInClassroom(c.Context(), classroomID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(users)

}


func (s *ClassroomService) GetAllClassrooms(c *fiber.Ctx) error {
	classrooms, err := s.store.GetAllClassrooms(c.Context())
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(classrooms)

}

func (s *ClassroomService) CreateClassroom(c *fiber.Ctx) error {

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

