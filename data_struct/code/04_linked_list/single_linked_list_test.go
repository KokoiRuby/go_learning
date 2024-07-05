package _4_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SingleLinkedListTestSuite struct {
	suite.Suite
	sll *SingleLinkedList
}

func (suite *SingleLinkedListTestSuite) SetupTest() {
	suite.sll = NewSingleLinkedList()
	for i := 0; i < 5; i++ {
		suite.sll.Append(i)
	}
	suite.sll.Print()
}

func (suite *SingleLinkedListTestSuite) TestInsertAfter() {
	ok, _ := suite.sll.InsertAfter(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.sll.Print()

	_, err := suite.sll.InsertAfter(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *SingleLinkedListTestSuite) TestInsertBefore() {
	ok, _ := suite.sll.InsertBefore(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.sll.Print()

	_, err := suite.sll.InsertBefore(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *SingleLinkedListTestSuite) TestRemove() {
	v := suite.sll.Remove(3)
	assert.Equal(suite.T(), 3, v)
	suite.sll.Print()
}

func TestArrayTestSuite(t *testing.T) {
	suite.Run(t, new(SingleLinkedListTestSuite))
}
