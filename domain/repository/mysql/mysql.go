package mongo

import (
	"context"
	"database/sql"
	"github.com/akuera/go-template/domain/models"
	"github.com/akuera/go-template/domain/repository"
)

type mySql struct {
	db *sql.DB
}

func (m mySql) CreateDomain(ctx context.Context, domain *models.Domain) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mySql) GetDomain(ctx context.Context, s string) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mySql) GetDomains(ctx context.Context) ([]*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mySql) UpdateDomain(ctx context.Context, s string, domain *models.Domain) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mySql) DeleteDomain(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func New(db *sql.DB) repository.Repo {
	return &mySql{
		db: db,
	}
}
