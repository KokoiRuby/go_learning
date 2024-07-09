package _9_bsearch

import (
	"errors"
)

func BSearch(a []int, v int, low, high int) (int, error) {
	if len(a) == 0 {
		return -1, errors.New("array is empty")
	}
	mid := low + ((high - low) >> 1)
	for low <= high {
		if a[mid] == v {
			return mid, nil
		} else if v > a[mid] {
			return BSearch(a, v, mid+1, high)
		} else {
			return BSearch(a, v, low, mid-1)
		}
	}
	return -1, errors.New("not found")
}

func BSearchFirst(a []int, v int, low, high int) (int, error) {
	if len(a) == 0 {
		return -1, errors.New("array is empty")
	}
	mid := low + ((high - low) >> 1)
	for low <= high {
		if a[mid] == v {
			if a[mid-1] != v {
				return mid, nil
			} else {
				return BSearchFirst(a, v, low, mid-1)
			}
		} else if v > a[mid] {
			return BSearchFirst(a, v, mid+1, high)
		} else {
			return BSearchFirst(a, v, low, mid-1)
		}
	}
	return -1, errors.New("not found")
}

func BSearchLast(a []int, v int, low, high int) (int, error) {
	if len(a) == 0 {
		return -1, errors.New("array is empty")
	}
	mid := low + ((high - low) >> 1)
	for low <= high {
		if a[mid] == v {
			if a[mid+1] != v {
				return mid, nil
			} else {
				return BSearchLast(a, v, mid+1, high)
			}
		} else if v > a[mid] {
			return BSearchLast(a, v, mid+1, high)
		} else {
			return BSearchLast(a, v, low, mid-1)
		}
	}
	return -1, errors.New("not found")
}

func BSearchFirstGT(a []int, v int, low, high int) (int, error) {
	if len(a) == 0 {
		return -1, errors.New("array is empty")
	}
	mid := low + ((high - low) >> 1)
	for low <= high {
		if a[mid] < v {
			return BSearchFirstGT(a, v, mid+1, high)
		} else {
			if a[mid-1] < v {
				return mid, nil
			} else {
				return BSearchFirstGT(a, v, low, mid-1)
			}
		}
	}
	return -1, errors.New("not found")
}

func BSearchLastLT(a []int, v int, low, high int) (int, error) {
	if len(a) == 0 {
		return -1, errors.New("array is empty")
	}
	mid := low + ((high - low) >> 1)
	for low <= high {
		if a[mid] > v {
			return BSearchLastLT(a, v, low, mid-1)
		} else {
			if a[mid+1] >= v {
				return mid, nil
			} else {
				return BSearchLastLT(a, v, mid+1, high)
			}
		}
	}
	return -1, errors.New("not found")
}
