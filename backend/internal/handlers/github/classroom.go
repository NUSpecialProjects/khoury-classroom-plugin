package github

import (
	"fmt"
	"net/http"
	"time"
    //"sort"

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

func (service *GitHubService) SyncStudentAssignments(c *fiber.Ctx) error {
	var syncData models.AssignmentSync
	err := c.BodyParser(&syncData)
	if err != nil {
		return err
	}
/*
	client, err := service.getClient(c)
	if err != nil {
		fmt.Println("SyncStudentAssignments - Failed to get Client", err)
		return err
	}
    

	// get all currently accepted student assignments from cr
	accepted_student_assignments, err := client.GetAcceptedAssignments(c.Context(), syncData.AssignmentID)
	if err != nil {
		fmt.Println("SyncStudentAssignments - Could not get classroom assignments")
		return err
	}

	// get list of student assignments currently in the db under this assignment
	studentAssmnts, err := service.store.GetStudentAssignmentByAssignment(c.Context(), syncData.AssignmentID)
	if err != nil {
        fmt.Println("SyncStudentAssignments - Could not get: ", err)
		return err
	}
    
    
    // Check every accepted assignment for a match in our db 
	for _, assignment := range accepted_student_assignments {
        var inDB = false


        // check every student assnment in our db
        for _, studentA := range studentAssmnts {
            var noMatch = false
            students, err := service.store.GetStudentAssignmentGroup(c.Context(), studentA.ID)
            
            if err != nil {
                fmt.Println("SyncStudentAssignments - Could not find group", err)
            } else {
                if len(students) == len(assignment.Students) {
                    var ghStudents []string
                    // Take the usernames out of the gh user objects
                    for _, s := range assignment.Students {
                        if (s.Login != nil) {
                            ghStudents = append(ghStudents, *s.Login)
                        }
                    }

                    // sort them for easy comparison
                    sort.Strings(ghStudents)
                    sort.Strings(students)

                    for i, _ := range ghStudents {
                        if ghStudents[i] != students[i] {
                            noMatch = true
                            break
                        }
                    }

                    if !noMatch {
                        inDB = true
                        break
                    }
        
                }

            }
        } // end of student assignment loop

        if !inDB {
            var newStudentAssignment models.StudentAssignment




        }


	}*/


	// store any student assignments that aren't already in the db

	return nil
}
