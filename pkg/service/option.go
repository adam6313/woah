package service

// options -
type options struct {
	// Addrs -
	Addrs string
}

// Option -
type Option func(*options)

// WithAdds -
func WithAdds(addrs string) Option {
	return func(opts *options) {
		opts.Addrs = addrs
	}
}
