package broadcast

import (
	"go-micro.dev/v4/broker"
)

// BroadCast -
type BroadCast interface {
	// Subscribe - 監聽
	Subscribe(topic string, fn broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error)

	// Publish - 推播
	Publish(topic string, header map[string]string, body interface{}, opts ...broker.PublishOption) error

	// OnSubscribe - 監聽
	OnSubscribe(topic string, i interface{}, fn Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error)

	// CloseAll -
	CloseAll()
}
