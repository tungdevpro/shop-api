package entity

import (
	"regexp"
	"strings"

	ValueError "github.com/tungdevpro/shop-api/constant/value_error"
	"github.com/tungdevpro/shop-api/helper"
)

type SignUpDto struct {
	FullName        string `json:"fullname" form:"fullname"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
}

func (s *SignUpDto) Validate() error {
	s.Email = strings.TrimSpace(s.Email)
	if err := helper.CheckWrongEmail(s.Email); err != nil {
		return err
	}

	s.Password = strings.TrimSpace(s.Password)
	if len(s.Password) < 6 {
		return ValueError.PasswordLength
	}

	s.PasswordConfirm = strings.TrimSpace(s.PasswordConfirm)
	if s.Password != s.PasswordConfirm {
		return ValueError.PasswordEqual
	}

	s.FullName = strings.TrimSpace(s.FullName)
	if len(s.FullName) != 0 {
		regex := regexp.MustCompile(`^[a-zA-Z\s]+$`)
		if err := regex.MatchString(s.FullName); !err {
			return ValueError.FullNameInvalid
		}
	}
	return nil
}
