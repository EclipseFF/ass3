package adapters

import "goPracitce/services/contact/internal/domain"

type contactUseCase interface {
	create() (domain.Contact, error)
	get() (domain.Contact, error)
	update() (domain.Contact, error)
	delete() error
}

type groupUseCase interface {
	create() (domain.Group, error)
	get() (domain.Contact, error)
	addToGroup() error
}
