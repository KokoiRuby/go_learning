package main

func main() {
	const (
		i = 1 << iota // 1<<0
		j = 3 << iota // 3<<1
		k             // 3<<2
		l             // 12<<3
	)
	// 1 6 12 24
	println(i, j, k, l)
}
