package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["key"] = "value"
	m["key"] = "value_new"

	// del by key
	delete(m, "key")

	// flush
	m = make(map[string]string)

	// if had k
	v, ok := m["k"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}

	// iter
	m["k1"] = "v1"
	m["k2"] = "v2"
	m["k3"] = "v3"
	for k, v := range m {
		println(k, v)
	}

	// complex
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "kokoi"
	studentMap["stu01"]["gender"] = "Male"
	studentMap["stu01"]["country"] = "CN"

	for k1, v1 := range studentMap {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Println(k2, v2)
		}
	}

}
