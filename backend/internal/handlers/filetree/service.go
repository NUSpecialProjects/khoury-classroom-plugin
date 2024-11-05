package filetree

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type FileTreeService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func newFileService(store storage.Storage, githubappclient github.GitHubAppClient) *FileTreeService {
	return &FileTreeService{store: store, githubappclient: githubappclient}
}
