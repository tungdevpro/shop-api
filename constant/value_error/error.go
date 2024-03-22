package valueerror

import "errors"

var (
	ErrCannotCreateUser      = errors.New("cannot create user")
	ErrEmailCannotBeBlank    = errors.New("email cannot be black")
	ErrPasswordCannotBeBlank = errors.New("password cannot be black")
	ErrPasswordLength        = errors.New("password must be at least 6 characters long")
	ErrPasswordEqual         = errors.New("password confirm must equal password")
	ErrUnauthorized          = errors.New("the account does not exist on the system")
	ErrFullNameInvalid       = errors.New("invalid string, contains other characters")
	ErrVerifiedYourAccount   = errors.New("you have not verified your account")
	ErrOtpLength             = errors.New("otp must be 5 characters long")
	ErrOtpNotEqual           = errors.New("otp does not match")
)
