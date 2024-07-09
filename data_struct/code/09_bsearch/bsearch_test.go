package _9_bsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BubbleSortTestSuite struct {
	suite.Suite
	a []int
}

func (suite *BubbleSortTestSuite) TestBSearch() {
	suite.a = []int{1, 3, 5, 6, 8}
	toFind := 6
	i, errs := BSearch(suite.a, toFind, 0, len(suite.a)-1)
	if errs != nil {
		assert.EqualError(suite.T(), errs, "not found")
	}
	assert.Equal(suite.T(), 3, i)
}

func (suite *BubbleSortTestSuite) TestBSearchFirst() {
	suite.a = []int{1, 3, 4, 5, 6, 8, 8, 8, 8, 18}
	toFind := 8
	i, errs := BSearchFirst(suite.a, toFind, 0, len(suite.a)-1)
	if errs != nil {
		assert.EqualError(suite.T(), errs, "not found")
	}
	assert.Equal(suite.T(), 5, i)
}

func (suite *BubbleSortTestSuite) TestBSearchLast() {
	suite.a = []int{1, 3, 4, 5, 6, 8, 8, 8, 8, 18}
	toFind := 8
	i, errs := BSearchLast(suite.a, toFind, 0, len(suite.a)-1)
	if errs != nil {
		assert.EqualError(suite.T(), errs, "not found")
	}
	assert.Equal(suite.T(), 8, i)
}

func (suite *BubbleSortTestSuite) TestBSearchFirstGT() {
	suite.a = []int{1, 2, 2, 2, 3, 4, 4, 6, 9, 9, 12}
	toFind := 3
	i, errs := BSearchFirstGT(suite.a, toFind, 0, len(suite.a)-1)
	if errs != nil {
		assert.EqualError(suite.T(), errs, "not found")
	}
	assert.Equal(suite.T(), 4, i)
}

func (suite *BubbleSortTestSuite) TestBSearchLastLT() {
	suite.a = []int{1, 2, 2, 2, 3, 4, 4, 6, 9, 9, 12}
	toFind := 3
	i, errs := BSearchLastLT(suite.a, toFind, 0, len(suite.a)-1)
	if errs != nil {
		assert.EqualError(suite.T(), errs, "not found")
	}
	assert.Equal(suite.T(), 3, i)
}

func TestBubbleSortTestSuite(t *testing.T) {
	suite.Run(t, new(BubbleSortTestSuite))
}
