package main

import (
	"fmt"
	"hash/fnv"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Key() string {
	return p.Name
}

func (p Person) Hash() uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(p.Key()))
	if err != nil {
		return 0
	}
	return h.Sum32()
}

func main() {
	people := make(map[Person]string)

	person1 := Person{Name: "Alice", Age: 25}
	person2 := Person{Name: "Bob", Age: 30}

	people[person1] = "Engineer"
	people[person2] = "Manager"

	lookupPerson := Person{Name: "Bob", Age: 30}
	job, found := people[lookupPerson]
	if found {
		fmt.Printf("%s's job is %s\n", lookupPerson.Name, job)
	} else {
		fmt.Printf("Job not found for %s\n", lookupPerson.Name)
	}
}
