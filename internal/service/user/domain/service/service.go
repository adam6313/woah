package service

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

// userActivityService -
type userService struct {
	node *snowflake.Node
	log  func(context.Context) *zap.SugaredLogger
}

// NewUserService-
func NewUserService(node *snowflake.Node) UserService {
	return &userService{
		node: node,
		log: func(ctx context.Context) *zap.SugaredLogger {
			return ctxzap.Extract(ctx).With(zap.String("entry", "UserService")).Sugar()
		},
	}
}
