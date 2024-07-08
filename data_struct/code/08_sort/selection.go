package _8_sort

import "errors"

func SelectionSort(a []int) error {
	l := len(a)
	if l == 0 {
		return errors.New("empty array")
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		// switch
		a[i], a[min] = a[min], a[i]
	}

	return nil
}
