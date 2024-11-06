package classrooms

import (
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Returns the classrooms the authenticated user is part of.
func (s *ClassroomService) getUserClassrooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Returns the details of a classroom.
func (s *ClassroomService) getClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		classroomData, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"classroom": classroomData,
		})
	}
}

// Creates a new classroom.
func (s *ClassroomService) createClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var classroomData models.Classroom
		err := c.BodyParser(&classroomData)
		if err != nil {
			return errs.InvalidRequestBody(models.Classroom{})
		}

		createdClassroom, err := s.store.CreateClassroom(c.Context(), classroomData)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"created_classroom": createdClassroom,
		})
	}
}

// Updates an existing classroom.
func (s *ClassroomService) updateClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		var classroomData models.Classroom
		error := c.BodyParser(&classroomData)
		if error != nil {
			return errs.InvalidRequestBody(models.Classroom{})
		}
		classroomData.ID = classroomID

		updatedClassroom, err := s.store.UpdateClassroom(c.Context(), classroomData)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"updated_classroom": updatedClassroom,
		})
	}
}

func (s *ClassroomService) updateClassroomName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		var classroomData models.Classroom
		error := c.BodyParser(&classroomData)
		if error != nil {
			return errs.InvalidRequestBody(models.Classroom{})
		}
		classroomData.ID = classroomID

		existingClassroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return err
		}
		existingClassroom.Name = classroomData.Name

		updatedClassroom, err := s.store.UpdateClassroom(c.Context(), existingClassroom)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"updated_classroom": updatedClassroom,
		})
	}
}

// Returns the users of a classroom.
func (s *ClassroomService) getClassroomUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		usersInClassroom, err := s.store.GetUsersInClassroom(c.Context(), classroomID)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"users_in_classroom": usersInClassroom,
		})
	}
}

// Removes a user from a classroom.
func (s *ClassroomService) removeUserFromClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Generates a token to join a classroom.
func (s *ClassroomService) generateClassroomToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
