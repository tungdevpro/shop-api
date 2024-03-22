package inmem

import "github.com/tungdevpro/shop-api/common"

type repository struct {
	baseCtx *common.BaseContext
}

func NewAuthRepository(baseCtx *common.BaseContext) *repository {
	return &repository{
		baseCtx: baseCtx,
	}
}
