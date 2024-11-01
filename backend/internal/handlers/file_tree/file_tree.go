package file_tree

import (
	"net/http"
	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/gofiber/fiber/v2"
)

func (s *FileTreeService) GetGitTree(c *fiber.Ctx) error {

	//Init params array
	paramNames := [2]string{"orgName", "repoName"}
	
	//Init params acc
	params := make(map[string]string)
	
	//Validate params
	for _, paramName := range paramNames{
		params[paramName] = c.Params(paramName)
		if params[paramName] == "" {
			//Throw API error
			return errs.MissingApiParamError(paramName)
		}

	}
	tree, err := s.githubappclient.GetGitTree(params["orgName"], params["repoName"])
	if err != nil {
		return errs.GithubIntegrationError(err)
	}
	return c.Status(http.StatusOK).JSON(tree)

}

func (s *FileTreeService) GetGitBlob(c *fiber.Ctx) error {
	//Init params array
	paramNames := [3]string{"orgName", "repoName", "sha"}
	
	//Init params acc
	params := make(map[string]string)
	
	//Validate params
	for _, paramName := range paramNames{
		params[paramName] = c.Params(paramName)
		if params[paramName] == "" {
			//Throw API error
			return errs.MissingApiParamError(paramName)
		}

	}

	
	//Make gh req
	content, err := s.githubappclient.GetGitBlob(params["orgName"], params["repoName"], params["sha"])
	if err != nil {
		//Throw Github error
		return errs.GithubIntegrationError(err)
	}
	return c.Status(http.StatusOK).Send(content)

}
