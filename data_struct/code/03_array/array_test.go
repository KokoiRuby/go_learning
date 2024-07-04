package _03_array

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/suite"
)

var capacity int = 5

type ArrayTestSuite struct {
	suite.Suite
	arr *Array
}

func (suite *ArrayTestSuite) SetupTest() {
	suite.arr = NewArray(5)
}

func (suite *ArrayTestSuite) TestInsert() {
	err := suite.arr.Insert(uint(7), 7)
	if err != nil {
		assert.EqualError(suite.T(), err, "index out of range")
	}

	for i := 0; i < capacity; i++ {
		err := suite.arr.Insert(uint(i), i)
		if err != nil {
			suite.T().Fatal(err.Error())
		}
	}

	assert.Equal(suite.T(), uint(5), suite.arr.Len())

	err = suite.arr.Insert(uint(3), 3)
	if err != nil {
		assert.EqualError(suite.T(), err, "full array")
	}

	_, err = suite.arr.Get(uint(7))
	if err != nil {
		assert.EqualError(suite.T(), err, "index out of range")
	}

}

func (suite *ArrayTestSuite) TestDelete() {
	for i := 0; i < capacity; i++ {
		err := suite.arr.Insert(uint(i), i)
		if err != nil {
			suite.T().Fatal(err.Error())
		}
	}

	_, err := suite.arr.Remove(uint(7))
	if err != nil {
		assert.EqualError(suite.T(), err, "index out of range")
	}

	val, err := suite.arr.Remove(uint(3))
	if err != nil {
		suite.T().Fatal(err.Error())
	}
	assert.Equal(suite.T(), 3, val)
	val, _ = suite.arr.Get(uint(3))
	assert.Equal(suite.T(), 0, val)

}

func TestArrayTestSuite(t *testing.T) {
	suite.Run(t, new(ArrayTestSuite))
}
