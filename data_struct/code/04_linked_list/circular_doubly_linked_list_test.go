package _4_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CircularDoublyLinkedListTestSuite struct {
	suite.Suite
	cdll *CircularDoublyLinkedList
}

func (suite *CircularDoublyLinkedListTestSuite) SetupTest() {
	suite.cdll = NewCircularDoublyLinkedList()
	for i := 0; i < 5; i++ {
		suite.cdll.Append(i)
	}
	suite.cdll.Print()
}

func (suite *CircularDoublyLinkedListTestSuite) TestInsertAfter() {
	ok, _ := suite.cdll.InsertAfter(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.cdll.Print()

	_, err := suite.cdll.InsertAfter(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *CircularDoublyLinkedListTestSuite) TestInsertBefore() {
	ok, _ := suite.cdll.InsertBefore(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.cdll.Print()

	_, err := suite.cdll.InsertBefore(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *CircularDoublyLinkedListTestSuite) TestRemove() {
	v := suite.cdll.Remove(3)
	assert.Equal(suite.T(), 3, v)
	suite.cdll.Print()
}

func TestCircularDoublyLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(CircularDoublyLinkedListTestSuite))
}
