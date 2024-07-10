package _12_tree

type BST struct {
	*BinaryTree
}

func NewBST(root *Node) *BST {
	return &BST{NewBinaryTree(root)}
}

func (bst *BST) Find(v interface{}) *Node {
	return nil
}

func (bst *BST) Insert(v interface{}) bool {
	return false
}

func (bst *BST) Remove(v interface{}) bool {
	return false
}
