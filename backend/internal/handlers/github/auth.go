package github

import (
	"context"
	"encoding/json"
	"io"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func githubLogin(c *fiber.Ctx, cfg config.GitHubClient) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := "randomstate" // actually make this random
	sess.Set("state", state)
	if err := sess.Save(); err != nil {
		return err
	}

	url := cfg.OAuthConfig().AuthCodeURL(state)
	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return nil
}

func githubCallback(c *fiber.Ctx, cfg config.GitHubClient) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}
	storedState := sess.Get("state")

	state := c.Query("state")
	if state != storedState {
		return c.SendString("States don't match!")
	}

	code := c.Query("code")

	oAuthConfig := cfg.OAuthConfig()

	token, err := oAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.SendString("Code-Token exchange failed")
	}

	client := oAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.SendString("User data fetch failed")
	}

	userData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON parsing failed")
	}

	var user map[string]interface{}
	if err := json.Unmarshal(userData, &user); err != nil {
		return c.SendString("JSON unmarshal failed")
	}

	sess.Set("user", user)
	sess.Set("token", token.AccessToken)
	if err := sess.Save(); err != nil {
		return err
	}

	return c.JSON(user)
}
