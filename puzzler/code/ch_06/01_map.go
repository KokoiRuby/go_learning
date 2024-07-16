package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m)
	//m["key"] = "value" // panic: assignment to entry in nil map
	m = make(map[string]string, 5)
	m["key"] = "value"
	fmt.Println(m)

	mm := map[string]string{
		"key": "value",
	}
	fmt.Println(mm)

	mmm := make(map[string]string)
	mmm["key"] = "value"

	// if had key
	v, ok := mm["key"]
	if ok {
		fmt.Println(v)
	}

}
