package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/w-omondi/budget-tracker.git/internal/services"
	"github.com/w-omondi/budget-tracker.git/internal/utils"
)

func AuthenticateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		responseUtil := utils.NewApiResponse[any]()
		authString := ctx.Get("Authorization")
		newAuthManager := services.NewFirebaseAuthManager()
		extracted, err := newAuthManager.ExtractTokenMetadata(authString)
		if err != nil {
			res := responseUtil.SendErrorResponse(err.Error(), err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(res)
		}

		token, err := newAuthManager.VerifyToken(extracted)
		if err != nil {
			res := responseUtil.SendErrorResponse(err.Error(), err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(res)
		}

		ctx.Locals("fuid", token.UID)
		return ctx.Next()
	}
}
