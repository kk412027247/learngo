package queue

type Queue []int

func (q *Queue) Push(v int){
	// 加了指针之后，就会修改引用的本身，也就是this 的实例被改变了
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	if len(*q) == 0 {
		return 0
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
