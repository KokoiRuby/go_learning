package _4_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SinglyLinkedListTestSuite struct {
	suite.Suite
	sll *SinglyLinkedList
}

func (suite *SinglyLinkedListTestSuite) SetupTest() {
	suite.sll = NewSinglyLinkedList()
	for i := 0; i < 5; i++ {
		suite.sll.Append(i)
	}
	suite.sll.Print()
}

func (suite *SinglyLinkedListTestSuite) TestInsertAfter() {
	ok, _ := suite.sll.InsertAfter(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.sll.Print()

	_, err := suite.sll.InsertAfter(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *SinglyLinkedListTestSuite) TestInsertBefore() {
	ok, _ := suite.sll.InsertBefore(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.sll.Print()

	_, err := suite.sll.InsertBefore(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *SinglyLinkedListTestSuite) TestRemove() {
	v := suite.sll.Remove(3)
	assert.Equal(suite.T(), 3, v)
	suite.sll.Print()
}

func TestArrayTestSuite(t *testing.T) {
	suite.Run(t, new(SinglyLinkedListTestSuite))
}
