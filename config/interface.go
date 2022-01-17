package config

import (
	"context"

	"go-micro.dev/v4/config"
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

	// Close - close connection and watch
	Close()

	// Watch - object is you want to watch services
	Watcher(ctx context.Context, object ...string) chan Values
}
