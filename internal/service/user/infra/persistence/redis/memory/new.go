package memory

import (
	"context"
	"woah/internal/service/user/domain/repository"
	"woah/internal/service/user/infra/config"

	goredis "github.com/go-redis/redis/v8"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"

	"go.uber.org/zap"
)

type repo struct {
	db     int
	client *goredis.Client
	log    func(ctx context.Context) *zap.SugaredLogger
}

// NewRepository -
func NewRepository(client *goredis.Client, c *config.Config) repository.RedisRepository {

	return &repo{
		client: client,
		log: func(ctx context.Context) *zap.SugaredLogger {
			return ctxzap.Extract(ctx).With(zap.String("entry", "redisRepository")).Sugar()
		},
	}
}
