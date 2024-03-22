package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tungdevpro/shop-api/common"
	ValueError "github.com/tungdevpro/shop-api/constant/value_error"
	"github.com/tungdevpro/shop-api/modules/auth"
	"github.com/tungdevpro/shop-api/modules/auth/entity"
)

type api struct {
	biz auth.Business
}

func NewApi(biz auth.Business) *api {
	return &api{biz: biz}
}

func (api *api) SignUpHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var dto entity.SignUpDto
		if err := c.BodyParser(&dto); err != nil {

			return c.JSON(common.NewErrorResp(nil, ValueError.ErrCannotCreateUser.Error(), ""))
		}

		_, err := api.biz.SignUp(c.Context(), &dto)
		if err != nil {
			return c.JSON(common.NewErrorResp(nil, ValueError.ErrCannotCreateUser.Error(), "Has error"))
		}

		return c.JSON(common.NewSuccessResp(dto, nil, nil, ""))

	}
}
