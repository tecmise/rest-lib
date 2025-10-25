package interceptor

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tecmise/rest-lib/pkg/headers"
)

func Auditory(c *fiber.Ctx) error {
	logrus.WithField("method", c.Method()).WithField("path", c.Path()).Debugf("Validando dados de auditoria")
	contextId := c.Get("Context-ID")
	var username string
	var userpool string
	if contextId == "" {
		contextId = uuid.Nil.String()
	}
	bearer, err := headers.GetBearerToken(c)
	if err != nil {
		logrus.Warnf("Error ao obter token: %s", err.Error())
	}
	if bearer != "" {
		result, decodeError := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
			logrus.Debugf("Iniciando parsing do token, header: %+v", token.Header)
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				logrus.Error("erro in signing method")
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}
			return token, nil
		})
		claims := result.Claims.(jwt.MapClaims)
		if claimUser, ok := claims["username"]; ok {
			username = fmt.Sprintf("%v", claimUser)
			logrus.Debugf("Username from token: %s", username)
		}
		if claimUser, ok := claims["userpool"]; ok {
			userpool = fmt.Sprintf("%v", claimUser)
			logrus.Debugf("Userpool from token: %s", username)
		}
		if decodeError != nil {
			logrus.Errorf("Error ao obter token: %s", decodeError.Error())
		}
	}

	goCtx := c.UserContext()
	newCtx := context.WithValue(goCtx, "context_id", contextId)
	if username != "" {
		newCtx = context.WithValue(newCtx, "user_id", username)
	} else {
		newCtx = context.WithValue(newCtx, "user_id", uuid.Nil.String())
	}

	if userpool != "" {
		newCtx = context.WithValue(newCtx, "userpool", userpool)
	} else {
		newCtx = context.WithValue(newCtx, "userpool", "unknown")
	}

	c.SetUserContext(newCtx)
	return c.Next()
}
