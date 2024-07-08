package _7_recursion

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FacTestSuite struct {
	suite.Suite
}

func (suite *FacTestSuite) TestFac1() {
	assert.Equal(suite.T(), 1, Fac1(0))
	assert.Equal(suite.T(), 1, Fac1(1))
	assert.Equal(suite.T(), 2, Fac1(2))
	assert.Equal(suite.T(), 6, Fac1(3))
	assert.Equal(suite.T(), 24, Fac1(4))
}

func (suite *FacTestSuite) TestFac2() {
	assert.Equal(suite.T(), 1, Fac2(0))
	assert.Equal(suite.T(), 1, Fac2(1))
	assert.Equal(suite.T(), 2, Fac2(2))
	assert.Equal(suite.T(), 6, Fac2(3))
	assert.Equal(suite.T(), 24, Fac2(4))
}

func TestFacTestSuite(t *testing.T) {
	suite.Run(t, new(FacTestSuite))
}
