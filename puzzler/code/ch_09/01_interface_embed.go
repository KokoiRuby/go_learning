package main

type A interface {
	test01()
}

type B interface {
	test02()
}

type C interface {
	A
	B
	test03()
}

type Stu struct {
}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

func main() {
	var stu Stu
	stu.test01()
	stu.test02()
	stu.test03()
}
