package config

import (
	"context"
	"encoding/json"
	"sync"

	mconfig "go-micro.dev/v4/config"
)

var once sync.Once

// root - implement
type root struct {
	// c -
	c mconfig.Config

	// Config -
	Config Config `json:"config"`

	// Services - services data
	Services interface{} `json:"services"`

	// Ch -
	Ch chan Values `json:"ch"`
}

type Values struct {
	// Target -
	Target string `json:"target"`

	// Value  -
	Value []byte `json:"value"`
}

// Config -
type Config struct {
	// Mod -
	Mod string `json:"mod"`

	// Log -
	Log string `json:"log"`
}

// Environment -
type CMD struct {
	// Run -
	Run string `json:"run"`
}

// Raw -
func (e CMD) Raw() []byte {
	c := &struct {
		CMD CMD `json:"cmd"`
	}{CMD: e}

	d, _ := json.Marshal(c)

	return d
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

// Watch -
func (r root) Watch(ctx context.Context, object ...string) chan Values {
	fn := func() {
		go watch(ctx, r.c, &r, object...)
	}

	// only once run
	once.Do(fn)

	return r.Ch
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
func (r root) Get(s ...string) []byte {
	return r.c.Get(s...).Bytes()
}
