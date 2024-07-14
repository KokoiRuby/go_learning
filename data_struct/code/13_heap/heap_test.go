package _12_heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HeapTestSuite struct {
	suite.Suite
	h *Heap
}

func (s *HeapTestSuite) SetupTest() {
	s.h = NewHeap(15)
	for i := 0; i < 14; i++ {
		s.h.Insert(i)
	}
	s.h.Print()
	s.T().Log('\n')
}

func (s *HeapTestSuite) TestInsert() {
	s.h.Insert(14)
	s.h.Print()
	s.T().Log('\n')
	s.h.Insert(15)
	s.h.Print()
	s.T().Log('\n')
}

func (s *HeapTestSuite) TestPop() {
	assert.Equal(s.T(), 13, s.h.Pop())
	s.h.Print()
	s.T().Log('\n')
}

func TestHeapTestSuite(t *testing.T) {
	suite.Run(t, new(HeapTestSuite))
}
