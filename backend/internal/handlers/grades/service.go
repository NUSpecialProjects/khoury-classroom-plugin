package grades

import (
  "github.com/CamPlume1/khoury-classroom/internal/storage"
  "github.com/CamPlume1/khoury-classroom/internal/config"
)


type GradeService struct {
  store storage.Storage
  userCfg *config.GitHubUserClient
}


func newGradeService(store storage.Storage, userCfg *config.GitHubUserClient) *GradeService {
  return &GradeService{store: store, userCfg: userCfg}

}
