package _11_lru

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type LRUCache1TestSuite struct {
	suite.Suite
	c *LRUCache1
}

func (s *LRUCache1TestSuite) SetupTest() {
	s.c = NewLRUCache1(5)
	for i := 0; i < 5; i++ {
		s.c.Put(strconv.Itoa(i), i)
	}
	s.c.Print()
	s.T().Log("\n")
}

func (s *LRUCache1TestSuite) TestPut() {
	s.c.Put(strconv.Itoa(5), 5)
	s.c.Print()
	s.T().Log("\n")
}

func (s *LRUCache1TestSuite) TestGet() {
	assert.Equal(s.T(), 2, s.c.Get(strconv.Itoa(2)))
	s.c.Print()
	s.T().Log("\n")
}

func (s *LRUCache1TestSuite) TestRemove() {
	s.c.Remove(strconv.Itoa(2))
	s.c.Print()
	s.T().Log("\n")
}

func TestLRUCache1TestSuite(t *testing.T) {
	suite.Run(t, new(LRUCache1TestSuite))
}
