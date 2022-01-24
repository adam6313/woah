package repository

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"
)

type UserRepository interface {
	// CreateUser - 建立使用者
	CreateUser(ctx context.Context, user *aggregate.User) error
}
