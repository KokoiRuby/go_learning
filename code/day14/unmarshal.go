package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	Name     string
	Age      int
	Birthday string //....
	Sal      float64
	Skill    string
}

func main() {
	str1 := "{\"Name\":\"Jack\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"Go Language\"}"
	var s S
	err1 := json.Unmarshal([]byte(str1), &s)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(s)

	str2 := "[{\"address\":\"NY\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"MX\",\"Hawaii\"],\"age\":\"20\",\"name\":\"tom\"}]"
	// no make needed
	var sl []map[string]interface{}
	err2 := json.Unmarshal([]byte(str2), &sl)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(sl)

}
