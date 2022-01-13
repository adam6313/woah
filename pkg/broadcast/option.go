package broadcast

import (
	"woah/pkg/encoder"

	"go-micro.dev/v4/broker"
)

// Option -
type Option func(*broadcast)

// WithEncoder -
func WithEncoder(encoder encoder.Encoder) Option {
	return func(b *broadcast) {
		b.coder = encoder
	}
}

// WithBroker -
func WithBroker(bk broker.Broker) Option {
	return func(b *broadcast) {
		b.broker = bk
	}
}
