package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 这个 node相当于this TreeNode相当于实例,并且现式地写一个实例
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// 函数都是传递值，如果要改变实例，需要传递了指针,改变实例
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Igonred")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
