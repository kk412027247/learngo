package container

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!"
	fmt.Println(s)
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))
	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 因为中文字符的存在，需要用[]rune 对字符串进行进行转换
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
