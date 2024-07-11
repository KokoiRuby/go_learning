package _6_queue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ArrayQueueTestSuite struct {
	suite.Suite
	q *ArrayQueue
}

func (suite *ArrayQueueTestSuite) SetupTest() {
	suite.q = NewArrayQueue(5)
	for i := 0; i < 5; i++ {
		err := suite.q.Enqueue(i)
		assert.Nil(suite.T(), err)
	}
}

func (suite *ArrayQueueTestSuite) TestEnqueue() {
	err := suite.q.Enqueue(5)
	assert.EqualError(suite.T(), err, "queue is full")

	v, _ := suite.q.Dequeue()
	assert.Equal(suite.T(), 0, v)

	err = suite.q.Enqueue(5)
	assert.Nil(suite.T(), err)
}

func (suite *ArrayQueueTestSuite) TestDequeue() {
	for i := 0; i < 5; i++ {
		v, _ := suite.q.Dequeue()
		assert.Equal(suite.T(), i, v)
	}
	_, err := suite.q.Dequeue()
	assert.EqualError(suite.T(), err, "queue is empty")
}

func TestArrayQueueTestSuite(t *testing.T) {
	suite.Run(t, new(ArrayQueueTestSuite))
}