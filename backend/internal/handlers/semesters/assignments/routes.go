package assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/semesters/file_tree"
	"github.com/CamPlume1/khoury-classroom/internal/handlers/semesters/student_assignments"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newAssignmentService(params.Store)

	protected := router.Group("/assignments")
	student_assignments.Routes(protected, params)
	file_tree.Routes(protected, params)

	protected.Get("", service.GetAllAssignments)

	protected.Post("", service.CreateAssignment)

	protected.Post("/rubrics", service.CreateRubric)

	protected.Post("/assignment", service.CreateStudentAssignment)

	protected.Post("/due_dates", service.CreateDueDate)

	protected.Post("/regrades", service.CreateRegrade)

}
