package command

import (
	"context"
	"fmt"
	"reflect"
)

type dispatch struct {
	m map[reflect.Type]content
}
type content struct {
	i  interface{}
	fn reflect.Value
}

// Dispatch -
type Dispatch interface {
	// Handle -
	Handle(ctx context.Context, cmd Command) error
}

// NewDispatch -
func NewDispatch(args ...interface{}) Dispatch {
	d := dispatch{
		m: make(map[reflect.Type]content),
	}

	for _, arg := range args {

		cbType := reflect.TypeOf(arg)
		//if cbType.Kind() != reflect.Struct {
		//panic("")
		//}

		var oPtr reflect.Value
		fmt.Println(cbType.Kind())
		if cbType.Kind() != reflect.Ptr {
			oPtr = reflect.New(cbType)
		} else {
			oPtr = reflect.New(cbType.Elem())
		}

		fmt.Println(oPtr)

	}
	//cbType := reflect.TypeOf(arg)
	//if cbType.Kind() != reflect.Struct {
	//panic("")
	//}

	//for i := 0; i < cbType.NumMethod(); i++ {
	//method := cbType.Method(i)

	//fnType := method.Func.Type()

	//numArgs := fnType.NumIn()
	//fmt.Println(numArgs)

	//argType := fnType.In(numArgs - 1)

	//fmt.Println(method.Name)
	//fmt.Println(method)
	//d.m[argType] = content{
	//i:  cbType,
	//fn: method.Func,
	//}
	//}
	//}

	return d
}

// Handle -
func (d dispatch) Handle(ctx context.Context, cmd Command) error {
	t, _ := cmd.Type()
	fmt.Println(t)

	v, ok := d.m[t]
	if !ok {
		panic("panic")
	}
	fmt.Println(v.fn)

	oV := []reflect.Value{
		reflect.ValueOf(v.i).Elem(),
		reflect.ValueOf(ctx),
		reflect.ValueOf(cmd),
	}

	v.fn.Call(oV)

	return nil
}

//case 3:
//subV := reflect.ValueOf(m.Subject)
//replyV := reflect.ValueOf(m.Reply)
//oV = []reflect.Value{subV, replyV, oPtr}

//nc, _ := nats.Connect(nats.DefaultURL)
//c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
//defer c.Close()

//// Simple Publisher
//c.Publish("foo", "Hello World")

//// Simple Async Subscriber
//c.Subscribe("foo", func(s string) {
//fmt.Printf("Received a message: %s\n", s)
//})
