package server

import (
	"context"
	"fmt"
	"github.com/akuera/go-template/domain/handlers"
	mongo2 "github.com/akuera/go-template/domain/repository/mongo"
	"github.com/akuera/go-template/domain/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func RunServer(cfg *Config) {
	r := gin.New()

	// Mongo Configurations
	mongoDB := setupMongo(cfg.Mongo)

	// Mongo Service Setups
	domainDB := mongo2.New(mongoDB)

	// Service setup
	domainsvc := service.New(domainDB)

	// Handler Setup
	handlers.AddDomainHandlers(r.Group("domain-ms"), domainsvc)

	err := r.Run(fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return
	}
}

func setupMongo(config MongoConfig) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s@%s",
		config.Scheme, config.Username, config.Password, config.Host)))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("group-ms")
}
