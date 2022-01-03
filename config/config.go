package config

import (
	mconfig "go-micro.dev/v4/config"
)

// root - implement
type root struct {
	// c -
	c mconfig.Config

	// Config -
	Config Config `json:"config"`

	// Services - services data
	Services interface{} `json:"services"`
}

// Config -
type Config struct {
	// Mod -
	Mod string `json:"mod"`

	// Log -
	Log string `json:"log"`
}

// AppName -
func (r root) AppName() string {
	return appName
}

// Conf -
func (r root) Conf() mconfig.Config {
	return r.c
}

// Close -
func (r root) Close() {
	r.c.Close()
}

// Mod -
func (r root) Mod() string {
	return r.Config.Mod
}

// Log -
func (r root) Log() string {
	return r.Config.Log
}

// Service - get service config
func (r root) Service(s string) []byte {
	return r.c.Get(services, s).Bytes()
}
