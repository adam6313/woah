package config

import (
	"context"

	"go-micro.dev/v4/config"

	hconf "github.com/tyr-tech-team/hawk/config"
)

// IConfig - config interface
type IConfig interface {
	// Conf -
	Conf() config.Config

	// AppName - get this app name
	AppName() string

	// Get - get config data
	Get(...string) []byte

	// Mod - get service mod
	Mod() string

	// Log - get log value
	Log() string

	// MongoConf - get mongo default config
	MongoConf() hconf.Mongo

	// Close - close connection and watch
	Close()

	// Watch - object is you want to watch services
	Watch(ctx context.Context, object ...string)
}
