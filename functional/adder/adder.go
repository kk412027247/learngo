package main

import "fmt"

// 这里返回的是一个闭包，包括被返回的函数，以及函数所链接到的所有变量。
// 所以这里的sum可以理解为一个局部变量，可以被累加
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	//a := adder()
	//for i := 0 ; i< 10; i++ {
	//	fmt.Printf("0 + 1 + 2 + ... + %d = %d\n", i, a(i))
	//}
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + 2 + ... + %d = %d\n", i, s)
	}
}
