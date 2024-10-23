package student_assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newStudentAssignmentService(params.Store, params.GitHubApp)

	protected := router.Group("/student-assignments")
	protected.Get("/:studentAssignmentID", service.GetStudentAssignment)
}
