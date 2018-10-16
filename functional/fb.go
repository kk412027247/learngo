package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 只要是个类型，就能实现接口
type intGen func() int

// 函数的接受者 写在前面，就可以用实例 加一个 . 来到调用，当然，放到括号后面阔也是可以的。
// 只是写法不同而已
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	// 用这个方法就可以获取到文本了，总之就可以了
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()

	printFileContents(f)

}
