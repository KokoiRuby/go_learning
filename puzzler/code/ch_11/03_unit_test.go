package main

import (
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	if result != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}

	result = Add(3, 3)
	if result != 6 {
		t.Errorf("Expected 3 + 3 to equal 6, but got %d", result)
	}

	result = Add(5, 5)
	if result != 10 {
		t.Log("Failing the test")
		t.Fail()
	}

	result = Add(1, 1)
	if result != 2 {
		t.Log("Failing the test and stopping execution")
		t.FailNow()
	}

	result = Add(4, 4)
	if result != 8 {
		t.Fatal("Expected 4 + 4 to equal 8")
	}

	result = Add(0, 0)
	if result != 0 {
		t.Fatalf("Expected 0 + 0 to equal 0, but got %d", result)
	}

	if true {
		t.Skip("Skipping this test")
	}

	t.Helper()
	if result != 0 {
		t.Fatal("Helper function used")
	}

	t.Log("This is a log message")

	t.Logf("This is a formatted log message: %d + %d = %d", 2, 2, 4)
}
