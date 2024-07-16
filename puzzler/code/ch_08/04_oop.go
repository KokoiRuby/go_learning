package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu *Student) say() string {
	infoStr := fmt.Sprintf("name: %s, gender: %s, age: %d, id: %d, score: %f",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
	return infoStr
}

func main() {
	stu1 := Student{
		name:   "Tom",
		gender: "Male",
		age:    20,
		id:     1,
		score:  1.0,
	}
	fmt.Println(stu1.say())

	var stu2 = Student{
		name: "Jack",
		age:  5,
	}
	fmt.Println(stu2.say())

	stu3 := Student{
		age:  18,
		name: "Mary",
	}
	fmt.Println(stu3.say())

	// stu4 := Student{"Bruce", 18}

	var stu5 *Student = &Student{
		name:   "Tom",
		gender: "Female",
	}
	fmt.Println(stu5.say())

	stu6 := &Student{
		name:   "Jack",
		gender: "Female",
	}
	fmt.Println(stu6.say())

}
