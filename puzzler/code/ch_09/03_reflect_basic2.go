package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectInterface(i interface{}) {
	rType := reflect.TypeOf(i)
	fmt.Println("rType:", rType)

	rVal := reflect.ValueOf(i)
	fmt.Println("rVal:", rVal)

	fmt.Printf("Type of rVal is: %T\n", rVal)

	n2 := rVal.Int() + 1
	fmt.Println("n2:", n2)

	// back to interface{}
	iv := rVal.Interface()
	// interface to type
	num2 := iv.(int)
	fmt.Println("num2:", num2)
	fmt.Println("iv:", iv)

}

func reflectStruct(i interface{}) {
	rType := reflect.TypeOf(i)
	fmt.Println("rtype:", rType)
	fmt.Printf("type of rType is %T\n", rType)

	rVal := reflect.ValueOf(i)
	fmt.Println("rVal:", rVal)
	fmt.Printf("type of rVal is %T\n", rVal)

	iv := rVal.Interface()
	fmt.Println("iv:", iv)
	fmt.Printf("type of iv is %T\n", iv)

	stu, ok := iv.(Student)
	if ok {
		fmt.Println(stu.Name)
	}

}

func main() {
	num := 100
	reflectInterface(num)

	stu := Student{
		Name: "Jack",
		Age:  18,
	}
	reflectStruct(stu)
}
