package _03_array

import (
	"errors"
	"fmt"
)

type Array struct {
	data   []int
	length uint
}

func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity), // lint:ignore S1019 arr slice needs initial capacity equals to capacity
		length: 0,
	}
}

func (arr *Array) Len() uint {
	return arr.length
}

func (arr *Array) isIndexOutOfRange(idx uint) bool {
	return idx >= uint(cap(arr.data))
}

func (arr *Array) Insert(idx uint, val int) error {
	if arr.Len() == uint(cap(arr.data)) {
		return errors.New("full array")
	}
	if idx != arr.Len() && arr.isIndexOutOfRange(idx) {
		return errors.New("index out of range")
	}
	//for i := arr.length; i > idx; i-- {
	//	arr.data[i] = arr.data[i-1]
	//}
	arr.data[idx] = val
	arr.length++
	return nil
}

func (arr *Array) Get(idx uint) (int, error) {
	if arr.isIndexOutOfRange(idx) {
		return 0, errors.New("index out of range")
	}
	return arr.data[idx], nil
}

func (arr *Array) Remove(idx uint) (int, error) {
	if arr.isIndexOutOfRange(idx) {
		return 0, errors.New("index out of range")
	}
	v := arr.data[idx]
	arr.data[idx] = 0
	arr.length--
	return v, nil
}

func (arr *Array) Print() {
	fmt.Println(arr.data)
}
