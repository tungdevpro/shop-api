package valueerror

import "errors"

var (
	CannotCreateUser      = errors.New("cannot create user")
	EmailCannotBeBlank    = errors.New("email cannot be black")
	PasswordCannotBeBlank = errors.New("password cannot be black")
	PasswordLength        = errors.New("password must be at least 6 characters long")
	PasswordEqual         = errors.New("password confirm must equal password")
	Unauthorized          = errors.New("the account does not exist on the system")
	FullNameInvalid       = errors.New("invalid string, contains other characters")
	VerifiedYourAccount   = errors.New("you have not verified your account")
	OtpLength             = errors.New("otp must be 5 characters long")
	OtpNotEqual           = errors.New("otp does not match")
)
