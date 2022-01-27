package item

import (
	"woah/internal/service/item/domain/repository"
	"woah/internal/service/item/infra/config"

	"go.mongodb.org/mongo-driver/mongo"
)

// C - collection
const (
	C = "item"
)

// repo -
type repo struct {
	db           *mongo.Client
	databaseName string
}

// New -
func New(m *mongo.Client, conf config.Config) repository.ItemRepo {
	return &repo{
		db:           m,
		databaseName: conf.Mongo.Database,
	}
}
