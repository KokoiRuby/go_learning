package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, Add(1, 2))
	assert.NotEqual(t, 4, Add(1, 2))
}

func TestPointer(t *testing.T) {
	assert.Nil(t, GetPointer(-1))
	assert.NotNil(t, GetPointer(1))
}

func TestAddTableDriven(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Add 1 and 2", 1, 2, 3},
		{"Add 3 and 4", 3, 4, 7},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, Add(tc.a, tc.b))
		})
	}
}
