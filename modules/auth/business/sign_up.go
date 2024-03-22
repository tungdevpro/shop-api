package business

import (
	"context"

	"github.com/tungdevpro/shop-api/modules/auth/entity"
)

func (business *business) SignUp(ctx context.Context, dto *entity.SignUpDto) (*entity.SignUpResponse, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	result, err := business.repo.SignUp(ctx, dto)
	return result, err
}
