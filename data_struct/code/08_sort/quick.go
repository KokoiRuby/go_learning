package _8_sort

import "errors"

func QuickSort(a []int) error {
	l := len(a)
	if l == 0 {
		return errors.New("empty array")
	} else if l == 1 {
		return nil
	} else if l == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return nil
	} else {
		i := partition(a)
		QuickSort(a[:i])
		QuickSort(a[i+1:])
		return nil
	}

}

// in-place
func partition(a []int) int {
	l := len(a)
	pivot := a[l-1]
	var i int
	for j := 0; j < l; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[l-1] = a[l-1], a[i]
	return i
}
