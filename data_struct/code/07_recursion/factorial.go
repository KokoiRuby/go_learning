package _7_recursion

func Fac1(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * Fac1(n-1)
	}
}

var m2 = make(map[int]int)

func Fac2(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		if v, ok := m2[n]; ok {
			return v
		}
		res := n * Fac2(n-1)
		m2[n] = res
		return res
	}
}
