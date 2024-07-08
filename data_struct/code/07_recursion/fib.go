package _7_recursion

func Fib1(n int) int {
	if n == 0 || n == 1 {
		return int(n)
	} else {
		return Fib1(n-1) + Fib1(n-2)
	}
}

var m1 = make(map[int]int)

func Fib2(n int) int {
	if n == 0 || n == 1 {
		return int(n)
	} else {
		// hit
		if v, ok := m1[n]; ok {
			return v
		}
		// miss → cal + save
		res := Fib2(n-1) + Fib2(n-2)
		m1[n] = res
		return res
	}
}
