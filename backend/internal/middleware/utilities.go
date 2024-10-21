package middleware

import (
	"errors"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/github/userclient"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/gofiber/fiber/v2"
)


/* Warning: Usage of Protected Middleware is a prerequisite to the use of this function */
func GetClient(c *fiber.Ctx, store storage.Storage, userCfg *config.GitHubUserClient) (*userclient.UserAPI, error) {
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		fmt.Println("FAILED TO GET USERID")
		return nil, errs.NewAPIError(500, errors.New("failed to retrieve userID from context"))
	}
	fmt.Println("UserID: ", userID)

	session, err := store.GetSession(c.Context(), userID)
	if err != nil {
		fmt.Println("FAILED TO GET SESSION", err)
		return nil, err
	}

	client, err := userclient.NewFromSession(userCfg.OAuthConfig(), &session)

	if err != nil {
		fmt.Println("FAILED TO CREATE CLIENT", err)
		return nil, err
	}

	fmt.Println("UserID: ", userID)
	fmt.Println("Client: ", client)
	return client, nil
}
