package response

import (
	"errors"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/tecmise/rest-lib/pkg/exceptions"
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

func AppErrorHandler(ctx *fiber.Ctx, err error) error {
	var appErr *exceptions.AppError
	if errors.As(err, &appErr) {

		if appErr.Type == exceptions.TypeInternal && appErr.OriginalErr != nil {
			log.Printf("Internal Server Error: %v\n", appErr.OriginalErr)
		}

		var statusCode int
		switch appErr.Type {
		case exceptions.TypeValidation, exceptions.TypeBusiness:
			statusCode = fiber.StatusBadRequest
			break
		case exceptions.TypeNotFound:
			statusCode = fiber.StatusNotFound
			break
		case exceptions.TypeInternal:
			statusCode = fiber.StatusInternalServerError
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"statusCode": fiber.StatusInternalServerError,
				"code":       "INT-01",
				"message":    "generic.error.message",
				"time":       time.Now(),
			})
		}

		fiberMap := fiber.Map{
			"statusCode": statusCode,
			"code":       appErr.Code,
			"message":    appErr.Code.GetProps().Message,
			"time":       time.Now(),
		}
		logrus.Infof("Error: %v", appErr.Message)

		return ctx.Status(statusCode).JSON(fiberMap)
	}

	// Verifique erros do próprio Fiber (ex: fiber.NewError)
	var e *fiber.Error
	if errors.As(err, &e) {
		return ctx.Status(e.Code).JSON(fiber.Map{
			"statusCode": e.Code,
			"code":       "INT-02",
			"message":    "generic.error.message",
			"time":       time.Now(),
		})
	}

	// Erro genérico (que não é um AppError nem fiber.Error)
	log.Printf("Unhandled Generic Error: %v\n", err)
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"statusCode": fiber.StatusInternalServerError,
		"code":       "INT-02",
		"message":    "generic.error.message",
		"time":       time.Now(),
	})
}
