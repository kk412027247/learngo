package queue

type Queue []interface{}


// interface{} 定义是范性
func (q *Queue) Push(v interface{}){
	// 加了指针之后，就会修改引用的本身，也就是this 的实例被改变了
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() interface{} {
	if len(*q) == 0 {
		return 0
	}
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
