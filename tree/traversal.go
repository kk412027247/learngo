package tree

import "fmt"

// 可以有两个重复的包名, 实质上是在同一个文件，只是拆开来些而已。
func (node *Node) Traverse() {
	node.TraverseFunc(func(_node *Node) {
		_node.Print()
	})
	fmt.Println()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
