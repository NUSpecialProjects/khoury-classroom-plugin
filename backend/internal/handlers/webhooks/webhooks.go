package webhooks

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	models "github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/github"
)

func (s *WebHookService) WebhookHandler(c *fiber.Ctx) error {
	var dispatch = map[string]func(c *fiber.Ctx) error{
		"pull_request":                s.PR,
		"pull_request_review_comment": s.PRComment,
		"pull_request_review_thread":  s.PRThread,
		"push":                        s.PushEvent,
	}
	event := c.Get("X-GitHub-Event", "")

	handler, exists := dispatch[event]
	if !exists {
		return c.SendStatus(400)
	}

	return handler(c)
}

func (s *WebHookService) PR(c *fiber.Ctx) error {
	println("PR webhook event")
	return c.SendStatus(200)
}

func (s *WebHookService) PRComment(c *fiber.Ctx) error {
	payload := models.PRComment{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	if payload.Comment.AuthorAssociation == "COLLABORATOR" {
		println("regrade request")
	}
	return c.SendStatus(200)
}

func (s *WebHookService) PRThread(c *fiber.Ctx) error {
	println("PR thread webhook event")
	return c.SendStatus(200)
}

func (s *WebHookService) PushEvent(c *fiber.Ctx) error {
	// Extract the 'payload' form value
	payload := c.FormValue("payload", "")
	if payload == "" {
		return errs.BadRequest(errors.New("missing payload"))
	}

	// Unmarshal the JSON payload into the PushEvent struct
	var pushEvent github.PushEvent
	err := json.Unmarshal([]byte(payload), &pushEvent)
	if err != nil {
		return errs.BadRequest(errors.New("invalid JSON payload"))
	}

	// Check if this is first commit in a repository
	isInitialCommit, err := s.isInitialCommit(pushEvent)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// If app bot triggered the initial commit, initialize the base repository
	if isInitialCommit && (pushEvent.Pusher == nil && *pushEvent.Pusher.Name == "khoury-classroom[bot]") {
		err = s.baseRepoInitialization(c, pushEvent)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *WebHookService) baseRepoInitialization(c *fiber.Ctx, pushEvent github.PushEvent) error {
	// Create development branch
	_, err := s.githubappclient.CreateBranch(c.Context(), *pushEvent.Repo.Organization,
		*pushEvent.Repo.Name,
		*pushEvent.Repo.MasterBranch,
		"dev")
	if err != nil {
		fmt.Println(err)
		return errs.InternalServerError()
	}

	return c.SendStatus(200)
}

func (s *WebHookService) isInitialCommit(pushEvent github.PushEvent) (bool, error) {
	if pushEvent.BaseRef == nil || pushEvent.Created == nil || pushEvent.GetBefore() == "" {
		return false, errs.BadRequest(errors.New("malformatted push event"))
	}

	isInitialCommit := pushEvent.BaseRef == nil && *pushEvent.Created &&
		pushEvent.GetBefore() == "0000000000000000000000000000000000000000"
	return isInitialCommit, nil
}
