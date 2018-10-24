package main

import "fmt"

func main() {
	slice := make([]string, 3, 5)
	fmt.Println(slice)

	slice1 := []string{"red","blue","green","yellow","pink"}
	fmt.Println(slice1)

	slice2 := []int{10, 20, 30}
	fmt.Println(slice2)


	// 以下三种都是创建一个nil的切片
	var slice3 []int
	fmt.Println(slice3)

	slice4 := make([]int,0)
	fmt.Println(slice4)

	slice5 := []int{}
	fmt.Println(slice5)

	slice6 := []int{10,20,30,40,50}
	fmt.Println( slice6)
	newSlice := slice6[1: 3]
	newSlice = append(newSlice, 60)
	// 由于共享底层数组的原因，原数组也被改变了
	fmt.Println(newSlice, slice6)

	slice7 := []int{10,20,30,40}
	newSlice = append(slice7, 50)
	fmt.Println(newSlice)

	source := []string{"apple","orange","plum","banana","grape"}

	slice8 := source[2:3:4]
	fmt.Println(slice8)

	slice9 := source[2:3:3]
	slice9 = append(slice9, "wtf")
	fmt.Println(slice9, source)



}
