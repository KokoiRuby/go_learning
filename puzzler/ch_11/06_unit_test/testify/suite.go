package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CalculatorTestSuite struct {
	suite.Suite
	StartingNumber int
}

// this function executes before the test suite begins execution
func (suite *CalculatorTestSuite) SetupSuite() {
	// set StartingNumber to one
	fmt.Println(">>> From SetupSuite")
	suite.StartingNumber = 1
}

// this function executes after all tests executed
func (suite *CalculatorTestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// this function executes before each test case
func (suite *CalculatorTestSuite) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("-- From SetupTest")
	suite.StartingNumber = 1
}

// this function executes after each test case
func (suite *CalculatorTestSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (suite *CalculatorTestSuite) TestAddOne() {
	fmt.Println("From TestAddOne")
	suite.StartingNumber += 1
	suite.Equal(2, suite.StartingNumber)
}

func (suite *CalculatorTestSuite) TestSubtractOne() {
	fmt.Println("From TestSubtractOne")
	suite.StartingNumber -= 1
	suite.Equal(0, suite.StartingNumber)
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
