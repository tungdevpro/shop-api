package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/tungdevpro/shop-api/modules/auth/entity"
)

type API interface {
	// LoginHandler() fiber.Handler
	SignUpHandler() fiber.Handler
	// VerifyOtpHandler() fiber.Handler
	// ForgotPwdHandler() fiber.Handler
	// ChangePwdHandler() fiber.Handler
}

type Business interface {
	SignUp(context.Context, *entity.SignUpDto) (*entity.SignUpResponse, error)
}

type Repository interface {
	SignUp(context.Context, *entity.SignUpDto) (*entity.SignUpResponse, error)
}
