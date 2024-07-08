package _6_queue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type LinkedListQueueTestSuite struct {
	suite.Suite
	q *LinkedListQueue
}

func (suite *LinkedListQueueTestSuite) SetupTest() {
	suite.q = NewLinkedListQueue(5)
	for i := 0; i < 5; i++ {
		err := suite.q.Enqueue(i)
		assert.Nil(suite.T(), err)
	}
	err := suite.q.Print()
	assert.Nil(suite.T(), err)
	suite.T().Logf("\n")
}

func (suite *LinkedListQueueTestSuite) TestEnqueue() {
	err := suite.q.Enqueue(5)
	assert.EqualError(suite.T(), err, "queue is full")

	v, _ := suite.q.Dequeue()
	assert.Equal(suite.T(), 0, v)

	err = suite.q.Enqueue(5)
	assert.Nil(suite.T(), err)

	err = suite.q.Print()
	assert.Nil(suite.T(), err)
	suite.T().Logf("\n")
}

func (suite *LinkedListQueueTestSuite) TestDequeue() {
	for i := 0; i < 5; i++ {
		v, _ := suite.q.Dequeue()
		assert.Equal(suite.T(), i, v)
	}
	_, err := suite.q.Dequeue()
	assert.EqualError(suite.T(), err, "queue is empty")
}

func TestLinkedListQueueTestSuite(t *testing.T) {
	suite.Run(t, new(LinkedListQueueTestSuite))
}
