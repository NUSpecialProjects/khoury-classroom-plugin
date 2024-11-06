package works

import (
	"fmt"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/gofiber/fiber/v2"
)

// Returns the student works for an assignment.
func (s *WorkService) getWorks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.Atoi(c.Params("classroom_id"))
		if err != nil {
			return errs.BadRequest(err)
		}
		assignmentID, err := strconv.Atoi(c.Params("assignment_id"))
		if err != nil {
			return errs.BadRequest(err)
		}

		works, err := s.store.GetWorks(c.Context(), classroomID, assignmentID)
		if err != nil {
			fmt.Print(err)
			return errs.InternalServerError()
		}
		return c.Status(200).JSON(fiber.Map{
			"student_works": works,
		})
	}
}

// Returns the details of a specific student work.
func (s *WorkService) getWork() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.Atoi(c.Params("classroom_id"))
		if err != nil {
			return errs.BadRequest(err)
		}
		assignmentID, err := strconv.Atoi(c.Params("assignment_id"))
		if err != nil {
			return errs.BadRequest(err)
		}
		studentWorkID, err := strconv.Atoi(c.Params("work_id"))
		if err != nil {
			return errs.BadRequest(err)
		}

		work, err := s.store.GetWork(c.Context(), classroomID, assignmentID, studentWorkID)
		if err != nil {
			fmt.Print(err)
			return errs.InternalServerError()
		}
		return c.Status(200).JSON(fiber.Map{
			"student_work": work,
		})
	}
}
