package service

import (
	"context"
	"log"
	"woah/config"
	"woah/service/user"
)

// options -
type options struct {
	ctx context.Context
	IC  config.IConfig
}

// Option -
type Option func(*options)

// WithIC -
func WithIC(ic config.IConfig) Option {
	return func(o *options) {
		o.IC = ic
	}
}

// IService -
type IService interface {
	// Run -
	Run()
}

// New -
func New(opts ...Option) IService {
	o := &options{
		ctx: context.Background(),
	}

	for _, fn := range opts {
		fn(o)
	}

	return o
}

// Run -
func (o *options) Run() {
	apply(o.ctx, o)
}

func apply(ctx context.Context, opts *options) {
	// get cmd with root config
	m, _ := opts.IC.Conf().Map()["cmd"].(map[string]interface{})

	// get run script
	r, ok := m["run"].(string)
	if !ok {
		log.Fatal("run script is fatal")
	}

	go user.Apply(ctx, r, opts.IC)
}
