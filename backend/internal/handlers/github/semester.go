package github

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (service *GitHubService) SyncAssignments(c *fiber.Ctx) error {
	// Get Assignments from GH Classroom
	var syncData models.ClassroomSync
	err := c.BodyParser(&syncData)
	if err != nil {
		return err
	}

	client, err := service.getClient(c)
	if err != nil {
		fmt.Println("SyncAssignments - Failed to get Client", err)
		return err
	}

	assignments, err := client.ListAssignmentsForClassroom(c.Context(), syncData.Classroom_id)
	if err != nil {
		fmt.Println("SyncAssignments - Could not get classroom assignments")
		return err
	}

	// Get ids of assignments in docsthe db
	assignment_classroom_ids, err := service.store.GetAssignmentIDs(c.Context())
	if err != nil {
		return err
	}

	// Any assignments not in our database should be added
	for _, assignment := range assignments {
		var inDB = false
		// Check if assignment is in our db
		for _, a_c_id := range assignment_classroom_ids {
			if assignment.ID == a_c_id.Assignment_Classroom_ID {
				inDB = true
				break
			}
		}

		// Add to db if not in it
		if !inDB {

			assignmentData := models.Assignment{}

			dueDate := assignment.Deadline
			// ensure assignment has a deadline
			if dueDate != nil {
				parsedTime, err := time.Parse(time.RFC3339, *dueDate)
				if err != nil {
					fmt.Println("SyncAssignments - error parsing time data", err)
				} else {
					assignmentData.MainDueDate = &parsedTime
				}
			}

			assignmentData.InsertedDate = time.Now()
			assignmentData.Assignment_Classroom_ID = assignment.ID
			assignmentData.Name = assignment.Title
			sem, err := service.store.GetSemesterByClassroomID(c.Context(), syncData.Classroom_id)
			if err != nil {
				fmt.Println("SyncAssignments - Failed to get classroom id: ", err)
			} else {
				assignmentData.ClassroomID = sem.ClassroomID
			}

			error := service.store.CreateAssignment(c.Context(), assignmentData)
			if error != nil {
				fmt.Println("SyncAssignments - Failed to add assignment to db", error)
			}

		}

	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Synced data",
	})
}

func (service *GitHubService) ListClassrooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		classrooms, err := client.GetUserClassrooms(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch classrooms"})
		}

		var assignments []models.ClassroomAssignment
		for _, classroom := range classrooms {
			assignments, err = client.ListAssignmentsForClassroom(c.Context(), classroom.ID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "failed to fetch assignments"})
			}
		}

		return c.Status(200).JSON(assignments)
	}
}

func (service *GitHubService) GetUserSemesters() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		orgs, err := client.GetUserOrgs(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch orgs"})
		}

		org_ids := []int64{}
		for _, org := range orgs {
			org_ids = append(org_ids, org.ID)
		}

		semesters, err := service.store.ListSemestersByOrgList(c.Context(), org_ids)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch semesters"})
		}

		active_semesters := []models.Semester{}
		inactive_semesters := []models.Semester{}
		for _, semester := range semesters {
			if semester.Active {
				active_semesters = append(active_semesters, semester)
			} else {
				inactive_semesters = append(inactive_semesters, semester)
			}
		}

		return c.Status(200).JSON(fiber.Map{
			"active_semesters":   active_semesters,
			"inactive_semesters": inactive_semesters,
		})
	}
}

func (service *GitHubService) CreateSemester() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		var requestBody struct {
			OrgID         int64  `json:"org_id"`
			ClassroomID   int64  `json:"classroom_id"`
			OrgName       string `json:"org_name"`
			ClassroomName string `json:"classroom_name"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		semester := models.Semester{
			OrgID:         requestBody.OrgID,
			ClassroomID:   requestBody.ClassroomID,
			OrgName:       requestBody.OrgName,
			ClassroomName: requestBody.ClassroomName,
			Active:        false,
		}

		semester, err = service.store.CreateSemester(c.Context(), semester)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create semester"})
		}

		err = client.CreateSemesterRoles(c.Context(), semester)
		if err != nil {
			service.store.DeleteSemester(c.Context(), semester.ClassroomID)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create semester roles"})
		}

		return c.Status(200).JSON(fiber.Map{"semester": semester})
	}
}

func (service *GitHubService) ActivateSemester() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid classroom id"})
		}

		var requestBody struct {
			Activate bool `json:"activate"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		semester := models.Semester{}
		err = nil

		if requestBody.Activate {
			semester, err = service.store.ActivateSemester(c.Context(), classroomID)
		} else {
			semester, err = service.store.DeactivateSemester(c.Context(), classroomID)
		}

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to modify semester"})
		}

		return c.Status(200).JSON(fiber.Map{"semester": semester})
	}
}
