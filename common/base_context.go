package common

import (
	"github.com/tungdevpro/shop-api/config"
	"gorm.io/gorm"
)

type BaseContext struct {
	Config   *config.Config
	Database *gorm.DB
}

func NewBaseContext(cfg *config.Config, db *gorm.DB) *BaseContext {
	base := BaseContext{
		Config:   cfg,
		Database: db,
	}

	return &base
}
