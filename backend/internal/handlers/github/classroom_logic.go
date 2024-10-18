package github

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

type Sync struct {
	Classroom_id int64 `json:"classroom_id"`
}

func (service *GitHubService) SyncAssignments(c *fiber.Ctx) error {
	fmt.Println("SyncAssignments - Begun sync attempt")
	// Get Assignments from GH Classroom
	var syncData Sync
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
	fmt.Println(assignments)   
	if err != nil {
		fmt.Println("SyncAssignments - Could not get classroom assignments")
		return err
	}

	// Get ids of assignments in the db
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
				fmt.Println("SyncAssignments - Found assignment in DB")
				inDB = true
				break
			}
		}

		// Add to db if not in it
		if !inDB {
			fmt.Println("SyncAssignments - Assignment not in DB, adding...")

      assignmentData := models.Assignment{}
			
      dueDate := assignment.Deadline
      // ensure assignment has a deadline
      if (dueDate != nil) {
			  parsedTime, err := time.Parse(time.RFC3339, *dueDate)
			  if err != nil {
			  	fmt.Println("SyncAssignments - error parsing time data", err)
			  	parsedTime = time.Now()
        }
			  assignmentData.MainDueDate = &parsedTime  
      }


      assignmentData.InsertedDate = time.Now()
      assignmentData.Assignment_Classroom_ID = assignment.ID
      assignmentData.Name = assignment.Title
      assignmentData.SemesterID = 1 // TODO: NEEDS TO NOT BE HARD CODED

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
