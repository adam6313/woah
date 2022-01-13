package broadcast

import (
	"go-micro.dev/v4/broker"
)

// Handler -
type Handler func(map[string]string, interface{}) error

// Subscribe -
func (b *broadcast) Subscribe(topic string, fn broker.Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	sub, err := b.broker.Subscribe(topic, fn, opts...)
	if err != nil {
		return nil, err
	}

	b.mux.Lock()
	defer b.mux.Unlock()

	b.subscriberList[topic] = sub

	return sub, nil
}

// OnSubscribe -
func (b *broadcast) OnSubscribe(topic string, v interface{}, fn Handler, opts ...broker.SubscribeOption) (broker.Subscriber, error) {
	sub, err := b.broker.Subscribe(topic, func(e broker.Event) error {
		switch v.(type) {
		case []byte, *[]byte:
		case string, *string:
		case error, *error:
		default:
			b.coder.Decode(e.Message().Body, v)
		}

		return fn(e.Message().Header, v)
	}, opts...)

	return sub, err
}

// Publish -
func (b *broadcast) Publish(topic string, header map[string]string, body interface{}, opts ...broker.PublishOption) error {
	data, err := b.coder.Encode(body)
	if err != nil {
		return err
	}

	e := &broker.Message{
		Header: header,
		Body:   data,
	}

	if err := b.broker.Publish(topic, e, opts...); err != nil {
		return err
	}

	return nil
}

// CloseAll -
func (b *broadcast) CloseAll() {
	b.mux.Lock()
	defer b.mux.Unlock()

	for _, v := range b.subscriberList {
		v.Unsubscribe()
	}
}
