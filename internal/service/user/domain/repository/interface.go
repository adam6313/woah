package repository

import (
	"context"
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
