package user

import (
	"context"
	"woah/internal/service/user/domain/repository"
	"woah/internal/service/user/infra/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	C = "user"
)

// Repository -
type repo struct {
	db           *mongo.Client
	databaseName string
	log          func(context.Context) *zap.SugaredLogger
	userRepo     repository.UserRepository
}

// NewRepository -
func NewRepository(m *mongo.Client, c *config.Config) repository.UserRepository {
	spew.Dump(c)
	return &repo{
		db:           m,
		databaseName: c.Mongo.Database,
		log: func(ctx context.Context) *zap.SugaredLogger {
			return ctxzap.Extract(ctx).With(zap.String("entry", "userRepository")).Sugar()
		},
	}
}
