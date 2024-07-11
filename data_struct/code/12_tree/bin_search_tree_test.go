package _12_tree

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BSTTestSuite struct {
	suite.Suite
	bst *BST
}

func (s *BSTTestSuite) SetupTest() {
	s.bst = NewBST(&Node{
		v: 33,
		left: &Node{
			v: 16,
			left: &Node{
				v: 13,
				right: &Node{
					v: 15,
				},
			},
			right: &Node{
				v: 18,
				left: &Node{
					v: 17,
				},
				right: &Node{
					v: 25,
					left: &Node{
						v: 19,
					},
					right: &Node{
						v: 27,
					},
				},
			},
		},
		right: &Node{
			v: 50,
			left: &Node{
				v: 34,
			},
			right: &Node{
				v: 58,
				left: &Node{
					v: 51,
				},
				right: &Node{
					v: 66,
				},
			},
		},
	})
}

func (s *BSTTestSuite) TestSearch() {
	found, isLeftOf, parent := s.bst.Search(19)
	assert.Equal(s.T(), 19, found.v)
	assert.Equal(s.T(), true, isLeftOf)
	assert.Equal(s.T(), 25, parent.v)
}

func (s *BSTTestSuite) TestInsert() {
	ok := s.bst.Insert(55)
	assert.Equal(s.T(), true, ok)
	found, isLeftOf, parent := s.bst.Search(19)
	assert.Equal(s.T(), 55, found.v)
	assert.Equal(s.T(), false, isLeftOf)
	assert.Equal(s.T(), 51, parent.v)
}

func (s *BSTTestSuite) TestRemove() {
	ok := s.bst.Insert(55)
	ok = s.bst.Remove(55)
	assert.Equal(s.T(), true, ok)
	found, isLeftOf, parent := s.bst.Search(55)
	assert.Nil(s.T(), found)
	assert.Equal(s.T(), false, isLeftOf)
	assert.Nil(s.T(), parent)

	ok = s.bst.Remove(13)
	assert.Equal(s.T(), true, ok)
	found, isLeftOf, parent = s.bst.Search(13)
	assert.Nil(s.T(), found)
	assert.Equal(s.T(), false, isLeftOf)
	assert.Nil(s.T(), parent)
	found, _, _ = s.bst.Search(16)
	assert.Equal(s.T(), 15, found.left.v)

	ok = s.bst.Remove(18)
	assert.Equal(s.T(), true, ok)
	found, isLeftOf, parent = s.bst.Search(18)
	assert.Nil(s.T(), found)
	assert.Equal(s.T(), false, isLeftOf)
	assert.Nil(s.T(), parent)
	found, _, _ = s.bst.Search(19)
	assert.Equal(s.T(), 17, found.left.v)
	assert.Equal(s.T(), 25, found.right.v)

}

func TestBSTTestSuite(t *testing.T) {
	suite.Run(t, new(BSTTestSuite))
}
