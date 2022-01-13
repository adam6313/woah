package broker

import (
	"go-micro.dev/v4/broker"
)

// NewMemoryBroker -
func NewMemoryBroker() (broker.Broker, error) {
	b := broker.NewMemoryBroker()

	if err := b.Connect(); err != nil {
		return nil, err
	}

	return b, nil
}
