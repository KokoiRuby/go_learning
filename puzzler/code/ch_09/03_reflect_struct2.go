package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Person interface {
	GetInfo() (string, error)
	SetInfo(string, int, float64) error
}

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func (s Student) GetInfo() (string, error) {
	var str = fmt.Sprintf("name:%s,age:%d,score:%f", s.Name, s.Age, s.Score)
	//获取结构体成员的值
	return str, nil
}

func (s *Student) SetInfo(name string, age int, score float64) error {
	s.Name = name
	s.Age = age
	s.Score = score
	return errors.New("set info error")
}

// print info
func printInfo(s Person) { //Person类型表示了一个接口类型，接口类型的值可以是任意类型的值，所以这里的s可以是任意类型的值
	t := reflect.TypeOf(s)                                               //获取s的动态类型
	v := reflect.ValueOf(s)                                              //获取s的动态值
	fmt.Println("动态类型：", t.Name(), "静态类型：", t.Kind().String())           //获取s的静态类型
	fmt.Println("动态值：", v, "静态值：", v.Elem())                             //获取s的静态值
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct { //判断s的动态型类是否是结构体类型
		fmt.Println("s is not struct")
		return
	}

	//获取结构体成员的数量
	fmt.Println("结构体成员的数量：", t.Elem().NumField())

	//获取结构体成员的信息
	for i := 0; i < t.Elem().NumField(); i++ {
		fmt.Println("结构体成员的信息：", t.Elem().Field(i))
	}

	//获取方法数量
	fmt.Println("方法数量：", t.Elem().NumMethod())

	//获取方法的参数数量
	fmt.Println("方法的参数数量：", t.Elem().Method(0).Type.NumIn())

	//获取方法的返回值数量
	fmt.Println("方法的返回值数量：", t.Elem().Method(0).Type.NumOut())

	//判断结构体有没有这个方法
	metadata, ok := t.Elem().MethodByName("GetInfo")
	if ok {
		fmt.Println("结构体有GetInfo方法")
	} else {
		fmt.Println("结构体有多少方法：", metadata)
	}

}

func main() {
	study := &Student{
		"张三",
		18,
		100,
	}
	var s Person = study
	s.SetInfo("李四", 20, 99.21)

	printInfo(s)

}
