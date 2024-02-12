package util

import "net/mail"

func ValidarEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
