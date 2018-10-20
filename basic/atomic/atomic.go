package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	// 可以直接加锁，但是不要轻易用这个锁
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("save increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()

}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
