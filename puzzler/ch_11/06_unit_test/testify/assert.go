package main

func Add(a int, b int) int {
	return a + b
}

func GetPointer(value int) *int {
	if value > 0 {
		return &value
	}
	return nil
}
