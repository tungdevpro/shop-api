package inmem

import (
	"context"

	"github.com/tungdevpro/shop-api/modules/auth/entity"

	UserEntity "github.com/tungdevpro/shop-api/modules/user/entity"

	ValueError "github.com/tungdevpro/shop-api/constant/value_error"
)

func (repo *repository) SignUp(ctx context.Context, dto *entity.SignUpDto) (*entity.SignUpResponse, error) {
	repo.baseContext.L.Lock()
	defer repo.baseContext.L.Unlock()

	db := repo.baseContext.GetInstanceDB()

	user := UserEntity.User{
		Email: dto.Email,
	}

	doc := db.Where(&user).First(&user)

	if doc.Error != nil || doc.RowsAffected == 0 {
		return nil, ValueError.Unauthorized
	}

	// if err := user.VerifyPassword(loginDto.Password); err != nil {
	// 	return nil, authEntity.ErrVerifiedYourAccount
	// }

	return nil, nil
}
