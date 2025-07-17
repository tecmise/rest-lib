package response

import (
	"github.com/gofiber/fiber/v2"
)

type (
	Result struct {
		Code    int         `json:"code"`
		Content interface{} `json:"content"`
	}
)

func NewError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
}

func NewBadRequestError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
}

func NewSuccess(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(value)
}

func NewCreated(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusCreated).JSON(value)
}

func NewAccepted(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusAccepted).JSON(value)
}

func NewNoContent(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}

func NewNotFound(ctx *fiber.Ctx, message string) error {
	if message != "" {
		return ctx.Status(fiber.StatusNotFound).JSON(message)
	} else {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
}

func NewUnprocessable(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(message)
}

func NewConflict(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusConflict).JSON(message)
}

func NewUnauthorizedError(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(message)
}
