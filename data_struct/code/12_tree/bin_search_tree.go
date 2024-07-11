package _12_tree

type BST struct {
	*BinaryTree
}

func NewBST(root *Node) *BST {
	return &BST{NewBinaryTree(root)}
}

func (bst *BST) Search(v interface{}) (found *Node, isLeftOf bool, parent *Node) {
	var p *Node
	node := bst.GetRoot()
	var isLeft bool

	for node != nil {
		nodeVal := node.v
		// base case
		if nodeVal == v {
			if larger(p, node) {
				isLeft = true
				return p, isLeft, node
			}
			isLeft = false
			return p, isLeft, node
		} else if larger(nodeVal, v) {
			// recursive case
			p = node
			node = node.right
		} else {
			// recursive case
			p = node
			node = node.left
		}
	}
	return nil, false, nil
}

func (bst *BST) Insert(v interface{}) bool {
	node := bst.GetRoot()
	for node != nil {
		nodeVal := node.v
		if nodeVal == v {
			return false
		} else if larger(nodeVal, v) {
			// base case
			if node.right != nil {
				node.right = NewNode(v)
				return true
			}
			// recursive case
			node = node.right
		} else {
			// base case
			if node.left != nil {
				node.left = NewNode(v)
				return true
			}
			// recursive case
			node = node.left
		}
	}
	return true
}

func (bst *BST) Remove(v interface{}) bool {
	found, isLeftOf, parent := bst.Search(v)
	if found == nil {
		return false
	}
	// no child
	if found.left == nil && found.right == nil {
		found = nil
		return true
	} else if found.right == nil && found.left != nil {
		// 1 child
		if isLeftOf {
			parent.left = found.left
			found = nil
			return true
		} else {
			parent.right = found.left
			found = nil
			return true
		}
	} else if found.left == nil && found.right != nil {
		// 1 child
		if isLeftOf {
			parent.left = found.right
			found = nil
			return true
		} else {
			parent.right = found.right
			found = nil
			return true
		}
	} else {
		// 2 children
		tmp := NewBST(found.right)
		minimum := tmp.Min()
		found.v = minimum.v
		minimum = nil
	}
	return false
}

func larger(a, b interface{}) bool {
	aInt := a.(int)
	bInt := b.(int)

	if aInt > bInt {
		return true
	} else {
		return false
	}
}

func (bst *BST) Min() *Node {
	node := bst.GetRoot()
	for node != nil {
		node = node.left
	}
	return node
}

func (bst *BST) Max() *Node {
	node := bst.GetRoot()
	for node != nil {
		node = node.right
	}
	return node
}
