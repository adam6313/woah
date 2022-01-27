package mongo

import (
	"woah/internal/service/item/infra/config"

	"github.com/tyr-tech-team/hawk/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewDial -
func NewDial(conf config.Config) (*mongo.Client, error) {
	return mongodb.NewDial(conf.Mongo)
}
