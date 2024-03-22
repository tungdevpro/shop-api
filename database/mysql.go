package database

import (
	"github.com/tungdevpro/shop-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateMysql(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.Dns), &gorm.Config{})
	return db, err
}
