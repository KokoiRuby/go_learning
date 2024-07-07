package _5_stack

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LinkedListStackTestSuite struct {
	suite.Suite
	s *LinkedListStack
}

func (suite *LinkedListStackTestSuite) SetupTest() {
	suite.s = NewLinkedListStack(5)
	for i := 0; i < 5; i++ {
		err := suite.s.Push(i)
		assert.Nil(suite.T(), err)
	}
}

func (suite *LinkedListStackTestSuite) TestPush() {
	err := suite.s.Push(5)
	assert.EqualError(suite.T(), err, "stack is full")
}

func (suite *LinkedListStackTestSuite) TestPop() {
	l := suite.s.length
	for i := 0; i < l; i++ {
		v, _ := suite.s.Pop()
		assert.Equal(suite.T(), l-i-1, v)
	}
	_, err := suite.s.Pop()
	assert.EqualError(suite.T(), err, "stack is empty")
}

func (suite *LinkedListStackTestSuite) TestFlush() {
	suite.s.Flush()
}

func TestLinkedListStack(t *testing.T) {
	suite.Run(t, new(LinkedListStackTestSuite))
}
