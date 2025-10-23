package validators

import (
	"github.com/gofiber/fiber/v2"
)

type ValidationOutputMiddleware[T Validatable] func(c *fiber.Ctx) error

// Deprecated: use github.com/tecmise/validation-lib instead
type Validatable interface {
}
