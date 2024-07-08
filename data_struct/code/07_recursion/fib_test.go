package _7_recursion

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FibTestSuite struct {
	suite.Suite
}

func (suite *FibTestSuite) TestFib1() {
	assert.Equal(suite.T(), 0, Fib1(0))
	assert.Equal(suite.T(), 1, Fib1(1))
	assert.Equal(suite.T(), 1, Fib1(2))
	assert.Equal(suite.T(), 2, Fib1(3))
	assert.Equal(suite.T(), 3, Fib1(4))
}

func (suite *FibTestSuite) TestFib2() {
	assert.Equal(suite.T(), 0, Fib2(0))
	assert.Equal(suite.T(), 1, Fib2(1))
	assert.Equal(suite.T(), 1, Fib2(2))
	assert.Equal(suite.T(), 2, Fib2(3))
	assert.Equal(suite.T(), 3, Fib2(4))
}

func TestFibTestSuite(t *testing.T) {
	suite.Run(t, new(FibTestSuite))
}
