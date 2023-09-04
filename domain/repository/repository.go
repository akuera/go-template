package repository

import (
	"context"
	"github.com/akuera/go-template/domain/models"
)

type Repo interface {
	CreateDomain(context.Context, *models.Domain) (*models.Domain, error)
	GetDomain(context.Context, string) (*models.Domain, error)
	GetDomains(ctx context.Context) ([]*models.Domain, error)
	UpdateDomain(context.Context, string, *models.Domain) (*models.Domain, error)
	DeleteDomain(context.Context, string) error
}
