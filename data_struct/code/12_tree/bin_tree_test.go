package _12_tree

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BinTreeTestSuite struct {
	suite.Suite
	t *BinaryTree
}

func (s *BinTreeTestSuite) SetupTest() {
	s.t = NewBinaryTree(&Node{
		v: 1,
		left: &Node{
			v: 2,
			left: &Node{
				v: 4,
			},
			right: &Node{
				v: 5,
			},
		},
		right: &Node{
			v: 3,
			left: &Node{
				v: 6,
			},
			right: &Node{
				v: 7,
			},
		},
	})
}

func (s *BinTreeTestSuite) TestPreOrderTraverse() {
	s.t.preOrderTraverse(s.t.GetRoot())
}

func (s *BinTreeTestSuite) TestInOrderTraverse() {
	s.t.inOrderTraverse(s.t.GetRoot())
}

func (s *BinTreeTestSuite) TestPostOrderTraverse() {
	s.t.postOrderTraverse(s.t.GetRoot())
}

func TestBinTreeTestSuite(t *testing.T) {
	suite.Run(t, new(BinTreeTestSuite))
}
