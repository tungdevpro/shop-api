package business

import (
	"context"
	"errors"

	"github.com/tungdevpro/shop-api/modules/auth/entity"
)

func (business *business) SignUp(ctx context.Context, dto *entity.SignUpDto) (*entity.SignUpResponse, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	return nil, errors.New("eroror")
}
