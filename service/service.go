package service

import (
	"context"
	"log"
	"woah/config"
	b "woah/pkg/broadcast"
	"woah/service/user"
)

// options -
type options struct {
	ctx       context.Context
	IC        config.IConfig
	broadcast b.Broadcast
}

// Option -
type Option func(*options)

// WithIC -
func WithIC(ic config.IConfig) Option {
	return func(o *options) {
		o.IC = ic
	}
}

// WithBroadcase -
func WithBroadcase(b b.Broadcast) Option {
	return func(o *options) {
		o.broadcast = b
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
	cmd, _ := opts.IC.Conf().Map()["cmd"].(map[string]interface{})

	// get run script
	run, ok := cmd["run"].(string)
	if !ok {
		log.Fatal("run script is fatal")
	}

	go user.Apply(ctx, run, opts.IC, opts.broadcast)
}
