package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "Goを勉強している"

	//
	fmt.Println(len(str))

	//
	for _, v := range []rune(str) {
		fmt.Printf("%c\n", v)
	}

	//
	fmt.Println(strconv.Atoi("666"))
	fmt.Println(strconv.Itoa(666))

	//
	b := []byte(str)
	fmt.Println(b)

	//
	s := string([]byte{77, 88, 99})
	fmt.Println(s)

	//
	var n int64 = 123
	fmt.Println(strconv.FormatInt(n, 2))
	fmt.Println(strconv.FormatInt(n, 8))
	fmt.Println(strconv.FormatInt(n, 10))
	fmt.Println(strconv.FormatInt(n, 16))

	//
	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Count(str, "Go"))

	//
	fmt.Println(str == "goを勉強している")
	fmt.Println(strings.EqualFold(str, "goを勉強している"))

	//
	fmt.Println(strings.Replace(str, "Go", "go", 1))

	//
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))

	//
	fmt.Println(strings.Split(str, "o"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, ","))

	//
	fmt.Println(strings.Trim(" haha ", " "))
	fmt.Println(strings.TrimLeft(" haha ", " "))
	fmt.Println(strings.TrimRight(" haha ", " "))
	fmt.Println(strings.TrimSpace(" haha "))
	fmt.Println(strings.TrimPrefix(str, "Go"))
	fmt.Println(strings.TrimSuffix(str, "Go"))

	//
	fmt.Println(strings.HasPrefix(str, "Go"))
	fmt.Println(strings.HasSuffix(str, "Go"))

}
