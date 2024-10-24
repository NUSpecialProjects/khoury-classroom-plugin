package assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/semesters/assignments/student_assignments"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newAssignmentService(params.Store)

	protected := router.Group("/assignments")

	protected.Get("", service.GetAssignmentsInSemester)
	protected.Post("", service.CreateAssignment)
	protected.Post("/rubrics", service.CreateRubric)
	protected.Post("/assignment", service.CreateStudentAssignment)
	protected.Post("/due_dates", service.CreateDueDate)
	protected.Post("/regrades", service.CreateRegrade)

	specific := router.Group("/assignments/:assignmentID")
	specific.Get("", service.GetAssignment)
	student_assignments.Routes(specific, params)

}
