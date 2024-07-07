package _5_stack

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArrayStackTestSuite struct {
	suite.Suite
	s *ArrayStack
}

func (suite *ArrayStackTestSuite) SetupTest() {
	suite.s = NewArrayStack(5)
	for i := 0; i < 5; i++ {
		err := suite.s.Push(i)
		assert.Nil(suite.T(), err)
	}
}

func (suite *ArrayStackTestSuite) TestPush() {
	err := suite.s.Push(5)
	assert.EqualError(suite.T(), err, "stack is full")
}

func (suite *ArrayStackTestSuite) TestPop() {
	l := len(suite.s.data)
	for i := 0; i < l; i++ {
		v, _ := suite.s.Pop()
		assert.Equal(suite.T(), l-i-1, v)
	}
	_, err := suite.s.Pop()
	assert.EqualError(suite.T(), err, "stack is empty")
}

func (suite *ArrayStackTestSuite) TestFlush() {
	suite.s.Flush()
	suite.s.Print()
}

func TestArrayStackTestSuite(t *testing.T) {
	suite.Run(t, new(ArrayStackTestSuite))
}
