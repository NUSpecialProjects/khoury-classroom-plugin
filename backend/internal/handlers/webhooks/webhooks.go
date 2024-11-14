package webhooks

import (
	"fmt"

	models "github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/github"
)

func (s *WebHookService) WebhookHandler(c *fiber.Ctx) error {
	var dispatch = map[string]func(c *fiber.Ctx) error{
		"pull_request":                s.PR,
		"pull_request_review_comment": s.PRComment,
		"pull_request_review_thread":  s.PRThread,
		"push":                        s.RepositoryCreation,
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

func (s *WebHookService) RepositoryCreation(c *fiber.Ctx) error {
	fmt.Println("Repository creation webhook event")

	// Parse the payload
	var pushEvent github.PushEvent
	err := c.BodyParser(&pushEvent)
	if err != nil {
		return err
	}

	// If this is not a base repository, ignore the event
	if !s.isInitialCommit(pushEvent) || *pushEvent.Pusher.Name != "khoury-classroom[bot]" {
		fmt.Println("Not a repository creation")
		return nil
	}

	fmt.Println("Repository created!")

	return c.SendStatus(200)
}

func (s *WebHookService) isInitialCommit(pushEvent github.PushEvent) bool {
	return pushEvent.BaseRef == nil && *pushEvent.Created &&
		pushEvent.GetBefore() == "0000000000000000000000000000000000000000"
}
