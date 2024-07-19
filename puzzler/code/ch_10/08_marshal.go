package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	Name     string
	Age      int
	Birthday string
	sal      float64
	skill    string
}

func main() {
	s := S{
		Name:     "Jack",
		Age:      20,
		Birthday: "2000-01-01",
		sal:      1000.0,
		skill:    "Go Language",
	}
	jsonStr1, err := json.Marshal(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr1))

	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "Jack"
	m["age"] = 18
	m["country"] = "US"

	jsonStr2, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr2))

	sl := []string{"Apple", "Banana", "Cherry"}
	jsonStr3, err := json.Marshal(&sl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr3))

	num := 123
	jsonStr4, err := json.Marshal(num)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr4))

}
