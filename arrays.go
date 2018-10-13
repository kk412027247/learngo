package main

import "fmt"

// 数组是值类型
func printArray(arr [5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}

}

// 通过传递数组指针，改变数组
func _printArray(arr *[5]int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

// 通过传递数组切片，改变数组
func printArray2(arr []int) {
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}

}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println("printArray1")
	printArray(arr1)
	fmt.Println("printArray3")
	printArray(arr3)
	fmt.Println("printArray1 3")
	fmt.Println(arr1, arr3)

	_printArray(&arr3)
	fmt.Println(arr3)

	printArray2(arr2[:])
	fmt.Println(arr3)
}
