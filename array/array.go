package main

import "fmt"

func main() {
	array := [5]*int{0:new(int), 1: new(int)}
	// 在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
	// 指针就是一个地址，在地址前面加一个* 就可以访问到地址的内容。
	// 变量前面加一个&就是取变量所在地址。
	// *&a == a
	a := 1
	fmt.Println(*&a == a)

	fmt.Println(*new(int))    // 0

	*array[0] = 10
	*array[1] = 20
	fmt.Println(*array[0], *array[1])

	var array1  [5]string
	array2 := [5]string{"red","blue","green","yellow","pink"}
	array1 = array2
	array2[2] = "233"
	fmt.Println(array1)
	fmt.Println(array2)

	var array3 [3]*string
	array4 := [3]*string{new(string), new(string), new(string)}

	*array4[0] = "red"
	*array4[1] = "blue"
	*array4[2] = "green"

	array3 = array4

	fmt.Println(array3,array4)

	for index,pointer := range array3{
		fmt.Println(index,*pointer)
	}




}
