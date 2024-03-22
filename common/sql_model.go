package common

import (
	"time"

	"github.com/tungdevpro/shop-api/helper"
)

type ID int64

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	Uuid      string     `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (sql *SQLModel) GenerateId() {
	sql.Uuid = helper.EncodeId(sql.Id)
}
