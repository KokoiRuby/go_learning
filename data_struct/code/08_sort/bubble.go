package _8_sort

import "errors"

func BubbleSort(a []int) error {
	l := len(a)
	if l == 0 {
		return errors.New("empty array")

	}
	for i := 0; i < l; i++ {
		// flag to exit outer loop
		for j := 0; j < l-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return nil
}
