package broadcast

import (
	"sync"
	"woah/pkg/encoder"
	"woah/pkg/encoder/json"

	"go-micro.dev/v4/broker"
)

// broker -
type broadcast struct {
	mux            *sync.Mutex
	coder          encoder.Encoder
	broker         broker.Broker
	subscriberList map[string]broker.Subscriber
}

// New -
func New(opts ...Option) BroadCast {
	bc := &broadcast{
		mux:            &sync.Mutex{},
		subscriberList: make(map[string]broker.Subscriber),
	}

	for _, fn := range opts {
		fn(bc)
	}

	return bc
}

func NewTest(b broker.Broker) BroadCast {
	j := json.NewEncoder()
	return New(
		WithEncoder(j),
		WithBroker(b),
	)
}
