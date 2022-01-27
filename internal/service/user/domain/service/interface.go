package service

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"

	"github.com/dgrijalva/jwt-go"
)

// UserService -
type UserService interface {
	// CheckUserData - 檢查使用者資料
	CheckUserData(ctx context.Context, user *aggregate.User) error

	// SetUpdateUser - 設置使用者更新內容
	SetUpdateUser(ctx context.Context, origin, in *aggregate.User) error

	// GenerateToken - 產生權杖
	GenerateToken(ctx context.Context, data jwt.Claims) (string, error)
}
