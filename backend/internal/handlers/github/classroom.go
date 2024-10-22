package github

import (
	"fmt"
	"net/http"
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
        } else {
			    assignmentData.MainDueDate = &parsedTime   
        }
      }


      assignmentData.InsertedDate = time.Now()
      assignmentData.Assignment_Classroom_ID = assignment.ID
      assignmentData.Name = assignment.Title
      sem, err := service.store.GetSemesterByClassroomID(c.Context(), syncData.Classroom_id)
      if (err != nil) {
        fmt.Println("SyncAssignments - Failed to get classroom id", err)
      } else {
        assignmentData.SemesterID = *sem.ID
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



func (s *GitHubService) GetAssignmentBy
