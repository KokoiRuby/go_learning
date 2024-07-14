package main

func main() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)
	// 0 1 2 ha ha 100 100 7 8
	println(a, b, c, d, e, f, g, h, i)
}
