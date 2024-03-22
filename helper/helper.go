package helper

import (
	"net/mail"
	"regexp"
)

func CheckWrongEmail(s string) error {
	_, err := mail.ParseAddress(s)

	return err
}

func HasDigit(s string) bool {
	regex := regexp.MustCompile(`\d`)

	return regex.MatchString(s)

}
