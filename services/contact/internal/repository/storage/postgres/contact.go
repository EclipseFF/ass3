package postgres

import (
	"context"
	"goPracitce/pkg/store/postgres"
	"goPracitce/services/contact/internal/domain"
)

type ContactRepository struct {
	pool *postgres.Store
}

func (r ContactRepository) CreateContact(ctx context.Context, contact *domain.Contact) error {
	return nil
}
