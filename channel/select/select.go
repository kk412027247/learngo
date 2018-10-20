package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			//time.Sleep(time.Second)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	// range 是等c传完之后，直接读这个数据,当 close 关闭channel时候使用
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d receiveed %d\n", id, n)
	}

}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	// 十秒钟之后退出
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int

		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		// channel 的条件选择可以用select, go都是通过这种方式进行同步的
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}

}
