package files

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type FileService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func newFileService(store storage.Storage, githubappclient github.GitHubAppClient) *FileService {
	return &FileService{store: store, githubappclient: githubappclient}
}
