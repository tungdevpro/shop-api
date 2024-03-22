package inmem

import "github.com/tungdevpro/shop-api/common"

type repository struct {
	baseContext *common.BaseContext
}

func NewAuthRepository(baseCtx *common.BaseContext) *repository {
	return &repository{
		baseContext: baseCtx,
	}
}
