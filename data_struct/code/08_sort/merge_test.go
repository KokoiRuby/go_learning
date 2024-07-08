package _8_sort

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MergeSortTestSuite struct {
	suite.Suite
	a []int
}

func (suite *MergeSortTestSuite) SetupTest() {
	// rand.Seed(time.Now().UnixNano())
	sl := make([]int, 10)
	for i := 0; i < 10; i++ {
		sl[i] = rand.Intn(10)
	}
	suite.a = sl
}

func (suite *MergeSortTestSuite) TestMergeSort() {
	err := MergeSort(suite.a)
	assert.Nil(suite.T(), err)
	for _, v := range suite.a {
		fmt.Printf("%v ", v)
	}
}

func TestMergeSortTestSuite(t *testing.T) {
	suite.Run(t, new(MergeSortTestSuite))
}
