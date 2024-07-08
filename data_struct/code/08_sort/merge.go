package _8_sort

import "errors"

func MergeSort(a []int) error {
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
		MergeSort(a[:l/2])
		MergeSort(a[l/2:])
		copy(a, merge(a[:l/2], a[l/2:]))
		return nil
	}
}

func merge(a, b []int) []int {
	len1 := len(a)
	len2 := len(b)
	tmp := make([]int, len1+len2)

	var i, j, k int // 0
	for ; i < len1 && j < len2; k++ {
		if a[i] <= b[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = b[j]
			j++
		}
	}

	// deal with left
	if i == len1 {
		for ; j < len2; j++ {
			tmp[k] = b[j]
			k++
		}
	} else {
		for ; i < len1; i++ {
			tmp[k] = a[i]
			k++
		}
	}

	return tmp
}
