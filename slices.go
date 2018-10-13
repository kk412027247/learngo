package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])

	s1 := arr[2:]
	fmt.Println(s1)

	s2 := arr[:]
	fmt.Println(s2)

	fmt.Println("after update slice1")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("after update slice2")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2

	s1 = arr[2:6]
	// 切片的切片 可以越界，只能向后扩展，不能向前扩展,不能超过底层数组的长度。
	s2 = s1[3:5]
	fmt.Println("arr = ", arr)
	//fmt.Println(s1[4])
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d \n", s2, len(s2), cap(s2))
	fmt.Println(s1[5:6])

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)

	// s4 s5 不再是arr 的view，添加元素如果超过cap，系统会重新分配一个 更大的底层数组。
	// append 必须有变量接受 返回值
	fmt.Println("arr = ", arr)
}
