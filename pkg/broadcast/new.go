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
func New() Broadcast {
	b := &broadcast{
		mux:            &sync.Mutex{},
		coder:          json.NewEncoder(),
		broker:         broker.NewMemoryBroker(),
		subscriberList: make(map[string]broker.Subscriber),
	}

	// connect (use local memory, so never have error)
	b.broker.Connect()

	return b
}
