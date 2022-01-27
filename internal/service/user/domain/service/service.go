package service

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

// userActivityService -
type userService struct {
	log func(context.Context) *zap.SugaredLogger
}

// NewUserService-
func NewUserService() UserService {
	return &userService{
		log: func(ctx context.Context) *zap.SugaredLogger {
			return ctxzap.Extract(ctx).With(zap.String("entry", "UserService")).Sugar()
		},
	}
}
