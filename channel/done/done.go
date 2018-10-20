package main

import (
	"fmt"
	"sync"
)

func doWork(id int, w worker) {
	// range 是等c传完之后，直接读这个数据,当 close 关闭channel时候使用
	for n := range w.in {
		fmt.Printf("worker %d receiveed %c\n", id, n)
		// 每次任务完成 声明一次任务完成
		w.done()
	}
}

type worker struct {
	in chan int
	// 因为是要影响外面的值，所以要用引用
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWork(id, w)
	return w
}

func chanDemo() {
	// 定义等待
	var wg sync.WaitGroup
	// 增加20个任务，因为一共有二十个
	wg.Add(20)

	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()

}
