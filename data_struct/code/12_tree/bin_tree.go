package _12_tree

import "fmt"

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(node *Node) *BinaryTree {
	return &BinaryTree{
		root: node,
	}
}

func (t *BinaryTree) GetRoot() *Node {
	return t.root
}

func (t *BinaryTree) preOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.String())
	t.preOrderTraverse(node.left)
	t.preOrderTraverse(node.right)
	return
}

func (t *BinaryTree) inOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	t.preOrderTraverse(node.left)
	fmt.Println(node.String())
	t.preOrderTraverse(node.right)
	return
}

func (t *BinaryTree) postOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	t.preOrderTraverse(node.left)
	t.preOrderTraverse(node.right)
	fmt.Println(node.String())
	return
}
