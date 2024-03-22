package entity

import (
	"github.com/tungdevpro/shop-api/common"
	"github.com/tungdevpro/shop-api/constant/table"
)

type User struct {
	*common.SQLModel `json:",inline"`
	Email            string `json:"email" gorm:"column:email;type:varchar(100);unique_index"`
	FullName         string `json:"fullname" gorm:"column:fullname;"`
	Password         string `gorm:"column:password;" json:"-"`
	AccessToken      string `json:"access_token" gorm:"column:access_token;"`
	IsEmailVerified  bool   `json:"is_email_verified" gorm:"column:is_email_verified;default:false"`
}

func (User) TableName() string { return table.USERS }
