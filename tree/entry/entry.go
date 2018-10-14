package main

import (
	"awesomeProject/tree"
	"fmt"
)


type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) posterOrder(){
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.posterOrder()
	right.posterOrder()

	myNode.node.Print()
}


// main 函数是 程序的入口，包名和文件夹名不一定需要一致，如果一包中有main，文件夹下面只能有一个main的文件
func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.posterOrder()
	fmt.Println()
}
