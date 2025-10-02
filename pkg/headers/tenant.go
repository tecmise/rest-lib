package headers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tecmise/rest-lib/pkg/keys"
	"github.com/tecmise/rest-lib/pkg/response"
)

func GetTenantId(ctx *fiber.Ctx) (string, error) {
	tenantId := ctx.Get(keys.HeaderXTenantID)
	if tenantId == "" {
		return "", response.NewUnauthorizedError(ctx, "X-Tenant-Id header is required")
	}

	return tenantId, nil
}
