package config

import (
	"go-micro.dev/v4/config"
)

// IConfig - config interface
type IConfig interface {
	// Conf -
	Conf() config.Config

	// AppName - get this app name
	AppName() string

	// Service - get service config data
	Service(string) []byte

	// Mod - get service mod
	Mod() string

	// Log - get log value
	Log() string

	// Close - close connection and watch
	Close()
}
