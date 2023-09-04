package mongo

import (
	"context"
	"github.com/akuera/go-template/domain/models"
	"github.com/akuera/go-template/domain/repository"
	"github.com/akuera/synon"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DomainsCollection = "domains"
)

type mgo struct {
	db *mongo.Database
}

func (m mgo) CreateDomain(ctx context.Context, domain *models.Domain) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mgo) GetDomain(ctx context.Context, s string) (*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mgo) GetDomains(ctx context.Context) ([]*models.Domain, error) {
	//TODO implement me
	panic("implement me")
}

func (m mgo) UpdateDomain(ctx context.Context, s string, domain *models.Domain) (*models.Domain, error) {
	fetchedDomain := models.Domain{}
	// Update with details
	updatedDomain := synon.Merge(fetchedDomain, domain).(*models.Domain)
	_, err := m.db.Collection(DomainsCollection).ReplaceOne(ctx, nil, updatedDomain)
	if err != nil {
		return nil, err
	}
	return updatedDomain, nil
}

func (m mgo) DeleteDomain(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func New(db *mongo.Database) repository.Repo {
	return &mgo{
		db: db,
	}
}
