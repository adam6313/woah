package mongo

import (
	"woah/internal/service/user/infra/config"

	"github.com/tyr-tech-team/hawk/infra/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewDial(c config.Config) (*mongo.Client, error) {

	return mongodb.NewDial(c.Mongo)
}
