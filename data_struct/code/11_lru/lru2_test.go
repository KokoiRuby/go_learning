package _11_lru

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LRUCacheTestSuite struct {
	suite.Suite
	c *LRUCache
}

func (s *LRUCacheTestSuite) SetupTest() {
	s.c = NewLRUCache(5)
	for i := 0; i < 5; i++ {
		s.c.Put(strconv.Itoa(i), i)
	}
	s.c.Print()
	s.T().Log("\n")
}

func (s *LRUCacheTestSuite) TestPut() {
	s.c.Put(strconv.Itoa(5), 5)
	s.c.Print()
	s.T().Log("\n")
}

func (s *LRUCacheTestSuite) TestGet() {
	assert.Equal(s.T(), 2, s.c.Get(strconv.Itoa(2)))
	s.c.Print()
	s.T().Log("\n")
}

func (s *LRUCacheTestSuite) TestRemove() {
	s.c.Remove(strconv.Itoa(2))
	s.c.Print()
	s.T().Log("\n")
}

func TestLRUCacheTestSuite(t *testing.T) {
	suite.Run(t, new(LRUCacheTestSuite))
}
