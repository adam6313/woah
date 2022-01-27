package repository

import (
	"context"
	"time"
	"woah/internal/service/user/domain/model/aggregate"
)

type UserRepository interface {
	// CreateUser - 建立使用者
	CreateUser(ctx context.Context, user *aggregate.User) error

	// UpdateUser - 更新使用者
	UpdateUser(ctx context.Context, user *aggregate.User) error

	// GetUser - 取得使用者
	GetUser(ctx context.Context, id string) (*aggregate.User, error)
}

type RedisRepository interface {
	// Get -
	Get(ctx context.Context, key string) (string, error)

	// Delete -
	Delete(ctx context.Context, key ...string) (int64, error)

	// SetNX -
	SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error)
}
