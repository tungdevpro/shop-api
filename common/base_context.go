package common

import (
	"sync"

	"github.com/tungdevpro/shop-api/config"
	"gorm.io/gorm"
)

type BaseContext struct {
	Config   *config.Config
	Database *gorm.DB
	L        *sync.RWMutex
}

func NewBaseContext(cfg *config.Config, db *gorm.DB) *BaseContext {
	base := BaseContext{
		Config:   cfg,
		Database: db,
		L:        new(sync.RWMutex),
	}

	return &base
}

func (bc *BaseContext) GetInstanceDB() *gorm.DB {
	return bc.Database.Session(&gorm.Session{})
}
