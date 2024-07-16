package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "Tom"
	rfv := reflect.ValueOf(&str)
	rfv.Elem().SetString("Jack")
	fmt.Println(str)

}
