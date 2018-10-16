package main

import (
	"bufio"
	"errors"
	"fmt"
	"learngo/functional/fib"
	"os"
)

// defer会生成一个栈，后进先出
// defer 会在函数即将退出的时候执行

func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occurred")
	//fmt.Println(4)
	for i := 0; i < 100; i++ {
		// defer 在语句时候计算，所以 是 30 ... 0 ,而不是打了一堆30出来。
		defer fmt.Println(i)
		if i == 30 {
			panic("printed tot many")
		}

	}

}

func writeFile(filename string) {
	//file, err := os.Create(filename)

	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a custom error")

	if err != nil {
		fmt.Println("Error:", err.Error())
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	// 建立了文件，就要关掉
	//defer file.Close()
	writer := bufio.NewWriter(file)
	//写进了buffer之后，需要写入文件。
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
	//tryDefer()
}
