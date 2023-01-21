package middleware

import (
	"encoding/base64"
	"github.com/RizkiMufrizal/belajar-gofiber/configuration"
	"github.com/RizkiMufrizal/belajar-gofiber/exception"
	"github.com/RizkiMufrizal/belajar-gofiber/model"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func BasicAuth(config configuration.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		username := config.Get("USERNAME")
		password := config.Get("PASSWORD")

		basicAuth := ctx.Get("Authorization")

		if basicAuth == "" {
			return ctx.
				Status(fiber.StatusBadRequest).
				JSON(model.GeneralResponse{
					Code:    404,
					Message: "Bad Request",
					Data:    "Header Not Found",
				})
		}

		basicAuthDecode, err := base64.StdEncoding.DecodeString(strings.Split(basicAuth, " ")[1])
		exception.PanicLogging(err)
		basicAuthDecodeString := string(basicAuthDecode)
		basicAuthUsername := strings.Split(basicAuthDecodeString, ":")[0]
		basicAuthPassword := strings.Split(basicAuthDecodeString, ":")[1]

		if username != basicAuthUsername && password != basicAuthPassword {
			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid Credentials",
				})
		}

		return ctx.Next()
	}
}
