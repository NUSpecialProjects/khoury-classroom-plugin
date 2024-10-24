package student_assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type StudentAssignmentService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func newStudentAssignmentService(store storage.Storage, githubappclient github.GitHubAppClient) *StudentAssignmentService {
	return &StudentAssignmentService{store: store, githubappclient: githubappclient}
}
