package submissions

import (
	"github.com/gofiber/fiber/v2"
)

// GetSubmissions returns the submissions for an assignment.
func (s *SubmissionService) GetSubmissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// GetSubmission returns the details of a submission.
func (s *SubmissionService) GetSubmission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// func (s *StudentAssignmentService) GetStudentAssignments(c *fiber.Ctx) error {
// 	classroomID, _ := strconv.ParseInt(c.Params("classroomID"), 10, 64)
// 	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)

// 	studentAssignment, err := s.store.GetStudentAssignments(c.Context(), classroomID, assignmentID)
// 	if err != nil {
// 		fmt.Println("error in service func")
// 		return err
// 	}

// 	return c.Status(http.StatusOK).JSON(studentAssignment)
// }

// func (s *StudentAssignmentService) GetStudentAssignment(c *fiber.Ctx) error {
// 	classroomID, _ := strconv.ParseInt(c.Params("classroomID"), 10, 64)
// 	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)
// 	studentAssignmentID, _ := strconv.ParseInt(c.Params("studentAssignmentID"), 10, 64)

// 	studentAssignment, err := s.store.GetStudentAssignment(c.Context(), classroomID, assignmentID, studentAssignmentID)
// 	if err != nil {
// 		fmt.Println("error in service func")
// 		return err
// 	}

// 	return c.Status(http.StatusOK).JSON(studentAssignment)
// }

// func (s *StudentAssignmentService) GetTotalStudentAssignments(c *fiber.Ctx) error {
// 	classroomID, _ := strconv.ParseInt(c.Params("classroomID"), 10, 64)
// 	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)

// 	total, err := s.store.GetTotalStudentAssignments(c.Context(), classroomID, assignmentID)
// 	if err != nil {
// 		fmt.Println("error in service func")
// 		return err
// 	}

// 	return c.Status(http.StatusOK).JSON(total)
// }
