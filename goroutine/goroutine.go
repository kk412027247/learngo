package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(_i int) {
			for {
				fmt.Printf("hello from goroutine %d\n", _i)
				//a[_i]++
				//// 主动交出线程
				//runtime.Gosched()
			}
		}(i)

	}
	// 如果不加sleep，会在 for 循环执行完之后，都来不及打印，所以要让进程等待一下
	time.Sleep(time.Minute)
	//fmt.Println(a)
}
