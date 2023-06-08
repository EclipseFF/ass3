package repository

import (
	"goPracitce/services/contact/internal/domain"
)

type ContactRepository interface {
	CreateContact() error
	GetContact() (domain.Contact, error)
}
