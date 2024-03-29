package create

import (
	"context"
)

// CreateUserUsecase -
type CreateUserUsecase interface {
	// Create - 建立使用者
	Create(ctx context.Context, command *CreateUser) (*UserCreated, error)
}
