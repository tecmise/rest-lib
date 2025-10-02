package headers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecmise/rest-lib/pkg/keys"
	"github.com/tecmise/rest-lib/pkg/response"
)

func GetXApiKey(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get(keys.HeaderXApiKey)
	if authHeader == "" {
		return "", response.NewUnauthorizedError(ctx, "x-api-ky header is required")
	}

	return authHeader, nil
}
