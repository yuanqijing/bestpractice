package adapter

import "sync"

type Interface interface {
}

func singletonWrapper(fn func() Interface) func() Interface {
	var instance Interface
	var once sync.Once
	return func() Interface {
		once.Do(func() {
			instance = fn()
		})
		return instance
	}
}

var setXXInstanceOnce sync.Once
var xXInstanceChan = make(chan Interface, 1)

func SetXXInstanceOnce(i Interface) {
	go setXXInstanceOnce.Do(func() {
		xXInstanceChan <- i
	})
}

var GetXXInstance = singletonWrapper(func() Interface {
	return <-xXInstanceChan
})
