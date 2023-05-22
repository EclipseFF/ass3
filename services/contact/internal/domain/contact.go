package domain

import (
	"github.com/pkg/errors"
	"strings"
)

type fullName struct {
	surname    string
	name       string
	patronymic string
}

type Contact struct {
	Id          uint
	FullName    fullName
	PhoneNumber string
}

func NewContact(id uint, surname, name, patronymic, phone string) (*Contact, error) {
	FIO := fullName{
		surname:    surname,
		name:       "name",
		patronymic: "patronymic",
	}

	contact := Contact{
		Id:          id,
		FullName:    FIO,
		PhoneNumber: phone,
	}

	if strings.ContainsAny(contact.PhoneNumber, "0123456789") {
		return nil, errors.New("phone number must not contain letters")
	}

	return &contact, nil

}
