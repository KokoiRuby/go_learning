package _12_tree

import "fmt"

/*
	starting from a[0]
	root=a[i]，left=a[2i+1]，right=a[2i+2]

	starting from a[1]
	root=a[i]，left=a[2i]，right=a[2i+1]
*/

type BinaryTreeArray struct {
	sl []interface{}
}

func NewBinaryTreeArray() *BinaryTreeArray {
	return &BinaryTreeArray{make([]interface{}, 0)}
}

func (t *BinaryTreeArray) GetSl() []interface{} {
	return t.sl
}

func (t *BinaryTreeArray) setSl(sl []interface{}) {
	t.sl = sl
}

func (t *BinaryTreeArray) preOrderTraverse(tree []interface{}, index int) {
	if index < len(tree) && tree[index] != nil {
		fmt.Printf("%v ", tree[index])
		t.preOrderTraverse(tree, 2*index+1)
		t.preOrderTraverse(tree, 2*index+2)
	}
}

func (t *BinaryTreeArray) inOrderTraverse(tree []interface{}, index int) {
	if index < len(tree) && tree[index] != nil {
		t.preOrderTraverse(tree, 2*index+1)
		fmt.Printf("%v ", tree[index])
		t.preOrderTraverse(tree, 2*index+2)
	}
}

func (t *BinaryTreeArray) postOrderTraverse(tree []interface{}, index int) {
	if index < len(tree) && tree[index] != nil {
		t.preOrderTraverse(tree, 2*index+1)
		t.preOrderTraverse(tree, 2*index+2)
		fmt.Printf("%v ", tree[index])
	}
}
