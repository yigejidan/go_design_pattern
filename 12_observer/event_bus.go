package _2_observer

import (
	"fmt"
	"reflect"
	"sync"
)

/*
通过观察者模式实现的事件总线，具备一下两个功能
	1、异步不阻塞
	2、支持任意参数值
*/

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// AsyncEventsBus 异步事件总线
type AsyncEventsBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventsBus {
	return &AsyncEventsBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (bus *AsyncEventsBus) Subscribe(topic string, f interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := bus.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handlers[topic] = handler

	return nil
}

func (bus *AsyncEventsBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}
