package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

// 这个 node相当于this treeNode相当于实例,并且现式地写一个实例
func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

// 函数都是传递值，如果要改变实例，需要传递了指针,改变实例
func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Igonred")
		return
	}
	node.value = value
}

func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//nodes := []treeNode {
	//	{value:3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)

	//root.print()
	root.right.left.setValue(4)
	//root.right.left.print()
	//fmt.Println()
	//
	//root.print()
	//root.setValue(100)

	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()

	//var pRoot *treeNode
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()

	root.traverse()
}
