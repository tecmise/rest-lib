package validators

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type ValidationOutputMiddleware[T Validatable] func(c *fiber.Ctx) error

type Validatable interface {
	Validate() error
}

type Toggleable interface {
	Validatable
	IsEnabled() bool
}

func JsonBinding(ctx *fiber.Ctx, input interface{}) error {
	// Faz o parsing do corpo para o input
	if err := ctx.BodyParser(&input); err != nil {
		return err
	}
	converted, ok := input.(Validatable)
	if ok {
		return converted.Validate()
	} else {
		return errors.New("error to convert input, is not validatable")
	}
}
