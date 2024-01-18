package util

import (
	"net/mail"

	"github.com/google/uuid"
)

func IsUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return err
	}
	return nil
}

func IsEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}
	return nil
}