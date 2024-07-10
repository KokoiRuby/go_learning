package _12_tree

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BinTreeByArrayTestSuite struct {
	suite.Suite
	t *BinaryTreeArray
}

func (s *BinTreeByArrayTestSuite) SetupTest() {
	s.t = NewBinaryTreeArray()
	s.t.setSl([]interface{}{1, 2, 3, 4, 5, 6, 7})
}

func (s *BinTreeByArrayTestSuite) TestPreOrderTraverse() {
	s.t.preOrderTraverse(s.t.GetSl(), 0)
}

func (s *BinTreeByArrayTestSuite) TestInOrderTraverse() {
	s.t.inOrderTraverse(s.t.GetSl(), 0)
}

func (s *BinTreeByArrayTestSuite) TestPostOrderTraverse() {
	s.t.postOrderTraverse(s.t.GetSl(), 0)
}

func TestBinTreeByArrayTestSuite(t *testing.T) {
	suite.Run(t, new(BinTreeByArrayTestSuite))
}
