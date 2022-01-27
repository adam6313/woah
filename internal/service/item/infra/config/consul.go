package config

import (
	"context"

	"github.com/tyr-tech-team/hawk/pkg/consul"
	"github.com/tyr-tech-team/hawk/srv"
)

// NewRegisterClient -
func NewRegisterClient(ctx context.Context) srv.Register {
	cc := consul.DefaultConsulConfig()
	cc.Address = C.Info.RemoteHost
	cli := consul.NewClient(ctx, cc)
	return cli
}
