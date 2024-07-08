package _8_sort

import "errors"

func InsertionSort(a []int) error {
	l := len(a)
	if l == 0 {
		return errors.New("empty array")
	}
	for i := 1; i < l; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				// 1 step forward
				a[j+1] = a[j]
			} else {
				// find the idx to insert
				break
			}
		}
		a[j+1] = v
	}
	return nil
}
