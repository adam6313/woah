package config

import (
	"context"

	"github.com/tyr-tech-team/hawk/config"
	"github.com/tyr-tech-team/hawk/config/source"
	"github.com/tyr-tech-team/hawk/pkg/consul"
	"github.com/tyr-tech-team/hawk/srv"
)

// C -
var C = Config{
	Info: config.Info{
		Name: "user",
	},
}

// Config -
type Config struct {
	Info    config.Info    `yaml:"info"`
	Mongo   config.Mongo   `yaml:"mongo"`
	Redis   config.Redis   `yaml:"redis"`
	Log     config.Log     `yaml:"log"`
	Service config.Service `yaml:"service"`
}

// NewConsulClient -
func NewConsulClient(ctx context.Context) consul.Client {
	cc := consul.DefaultConsulConfig()
	cc.Address = C.Info.RemoteHost
	cli := consul.NewClient(ctx, cc)
	return cli
}

// RemoteConfig -
func RemoteConfig(cli consul.Client) (Config, error) {
	r := config.NewReader(source.NewConsul(cli, C.Info.Name), config.YAML)
	err := r.ReadWith(&C)

	return C, err
}

// RegisterClient -
func RegisterClient(cli consul.Client) srv.Register {
	return cli
}
