package _12_tree

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BSTTestSuite struct {
	suite.Suite
	t *BST
}

func (s *BSTTestSuite) SetupTest() {

}

func TestBSTTestSuite(t *testing.T) {
	suite.Run(t, new(BSTTestSuite))
}
