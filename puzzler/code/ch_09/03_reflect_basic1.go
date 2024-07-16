package main

import (
	"fmt"
	"reflect"
)

type myInt int

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) Hello(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	fmt.Println()
}

func main() {
	u := &User{1, "Jack", 18}
	u.Hello(u)

	u.Hello(123)
	u.Hello(123.321)

	arr := [3]int{1, 2, 3}
	u.Hello(arr)

	sl := []int{3, 2, 1}
	u.Hello(sl)

	m := map[string]int{"key1": 1, "key2": 2}
	u.Hello(m)
}
