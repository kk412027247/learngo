package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// go中函数传递参数 都是值传递
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s" + op)
	}
}

func div(a, b int) (q, r int) {
	//return a / b, a % b
	q = a / b
	r = a % b
	return q, r
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func _swap(a, b *int) {
	*a, *b = *b, *a
}

func some(num ...int) {
	for _, v := range num {
		fmt.Println(v)
	}
}

func main() {

	some(1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 1)

	if result, err := eval(1, 2, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	if result, err := eval(1, 2, "+"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)

	c, d := 2, 8
	_swap(&c, &d)
	fmt.Println(c, d)
}
