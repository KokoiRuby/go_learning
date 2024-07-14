package main

import "fmt"

func main() {
	str1 := "big"
	b := []byte(str1)
	b[0] = 'p'
	fmt.Println(string(b))

	str2 := "pig"
	r := []rune(str2)
	r[0] = 'b'
	fmt.Println(string(r))
}
