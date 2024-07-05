package _4_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DoublyLinkedListTestSuite struct {
	suite.Suite
	dll *DoublyLinkedList
}

func (suite *DoublyLinkedListTestSuite) SetupTest() {
	suite.dll = NewDoublyLinkedList()
	for i := 0; i < 5; i++ {
		suite.dll.Append(i)
	}
	suite.dll.Print()
}

func (suite *DoublyLinkedListTestSuite) TestInsertAfter() {
	ok, _ := suite.dll.InsertAfter(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.dll.Print()

	_, err := suite.dll.InsertAfter(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *DoublyLinkedListTestSuite) TestInsertBefore() {
	ok, _ := suite.dll.InsertBefore(3, 33)
	assert.Equal(suite.T(), true, ok)
	suite.dll.Print()

	_, err := suite.dll.InsertBefore(100, 1000)
	assert.EqualError(suite.T(), err, "out of range")
}

func (suite *DoublyLinkedListTestSuite) TestRemove() {
	v := suite.dll.Remove(3)
	assert.Equal(suite.T(), 3, v)
	suite.dll.Print()
}

func TestDoublyLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(DoublyLinkedListTestSuite))
}
