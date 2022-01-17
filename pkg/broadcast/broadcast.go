package broadcast

import (
	"errors"
	"reflect"

	"go-micro.dev/v4/broker"
)

// Handler -
//type Handler func(map[string]string, interface{}) error

// Handler -
type Handler interface{}

var emptyType = reflect.TypeOf(&broker.Message{})

// Sub -
func (b *broadcast) Sub(topic string, cb Handler) {
	// cb is required
	if cb == nil {
		return
	}

	// get arguments type and num
	argType, numArgs, err := argInfo(cb)
	if argType == nil || err != nil {
		return
	}

	// reflect cb to Value
	cbValue := reflect.ValueOf(cb)

	fn := func(e broker.Event) error {
		var oV []reflect.Value

		switch argType {
		case emptyType:
			oV = []reflect.Value{reflect.ValueOf(e)}
		default:
			var oPtr reflect.Value
			if argType.Kind() != reflect.Ptr {
				oPtr = reflect.New(argType)
			} else {
				oPtr = reflect.New(argType.Elem())
			}

			// decoder
			b.coder.Decode(e.Message().Body, oPtr.Interface())
			if argType.Kind() != reflect.Ptr {
				oPtr = reflect.Indirect(oPtr)
			}

			// Callback
			switch numArgs {
			case 1:
				oV = []reflect.Value{oPtr}
			}

		}

		cbValue.Call(oV)

		return nil
	}

	sub, _ := b.broker.Subscribe(topic, fn)

	// lock
	b.mux.Lock()
	defer b.mux.Unlock()

	// set map
	b.subscriberList[topic] = sub
}

// Pub -
func (b *broadcast) Pub(topic string, v interface{}) error {
	data, err := b.coder.Encode(v)
	if err != nil {
		return err
	}

	e := &broker.Message{
		Body: data,
	}

	return b.broker.Publish(topic, e)
}

// CloseAll -
func (b *broadcast) CloseAll() {
	b.mux.Lock()
	defer b.mux.Unlock()

	for _, v := range b.subscriberList {
		v.Unsubscribe()
	}
}

func argInfo(cb Handler) (reflect.Type, int, error) {
	cbType := reflect.TypeOf(cb)
	if cbType.Kind() != reflect.Func {
		return nil, 0, errors.New("Handler need to be a func")
	}

	numArgs := cbType.NumIn()
	if numArgs == 0 {
		return nil, numArgs, nil
	}

	return cbType.In(numArgs - 1), numArgs, nil
}
