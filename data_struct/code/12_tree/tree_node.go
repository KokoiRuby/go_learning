package _12_tree

import "fmt"

type Node struct {
	v     interface{}
	left  *Node
	right *Node
}

func NewNode(d interface{}) *Node {
	return &Node{v: d}
}

func (node *Node) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", node.v, node.left, node.right)
}
