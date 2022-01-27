package redis

import (
	"woah/internal/service/user/infra/config"

	goredis "github.com/go-redis/redis/v8"
	"github.com/tyr-tech-team/hawk/infra/redis"
)

// NewDial - Redis連線
func NewDial(c *config.Config) (*goredis.Client, error) {
	cli, err := redis.NewDial(c.Redis)

	return cli, err
}
