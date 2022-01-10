package service

import "context"

// IService -
type IService interface {
	// Run -
	//Run(ctx context.Context)
}

// New -
func New(ctx context.Context, opts ...Option) IService {
	o := &options{}

	for _, fn := range opts {
		fn(o)
	}

	return o
}
