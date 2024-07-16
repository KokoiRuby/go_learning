package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
}

func (stu Student) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (stu Student) Set(name string, age int, score float32) {
	stu.Name = name
	stu.Age = age
	stu.Score = score
}

func (stu Student) Print() {
	fmt.Println("---start~----")
	fmt.Println(stu)
	fmt.Println("---end~----")
}

func reflectStruct(i interface{}) {
	t := reflect.TypeOf(i)
	val := reflect.ValueOf(i)
	k := val.Kind()
	if k != reflect.Struct {
		fmt.Println("Expect struct type, exiting...")
		return
	}

	fieldNum := val.NumField()
	fmt.Println(fieldNum)

	for i := 0; i < fieldNum; i++ {
		fmt.Println(val.Field(i))
		tagVal := t.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Println(tagVal)
		}
	}

	methodNum := val.NumMethod()
	fmt.Println(methodNum)

	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Method(0).Call(params)
	fmt.Println(res[0].Int())
}

func main() {
	stu := Student{
		Name:  "Jack",
		Age:   18,
		Score: 89.0,
	}
	reflectStruct(stu)
}
