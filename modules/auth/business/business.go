package business

import (
	"github.com/tungdevpro/shop-api/modules/auth"
)

type business struct {
	repo auth.Repository
}

func NewBusiness(repo auth.Repository) *business {
	return &business{
		repo: repo,
	}
}
