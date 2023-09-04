package service

import (
	"context"
	"github.com/akuera/go-template/domain"
	"github.com/akuera/go-template/domain/models"
	"github.com/akuera/go-template/domain/repository"
)

type service struct {
	repo repository.Repo
}

func (s service) CreateDomain(ctx context.Context, m *models.Domain) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetDomain(ctx context.Context, s2 string) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetDomains(ctx context.Context) ([]*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) UpdateDomain(ctx context.Context, s2 string, m *models.Domain) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) DeleteDomain(ctx context.Context, s2 string) error {
	//TODO implement me
	panic("implement me")
}

func New(repo repository.Repo) domain.Int {
	return &service{
		repo: repo,
	}
}
