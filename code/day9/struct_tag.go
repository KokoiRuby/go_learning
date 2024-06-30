package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Dog struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Color string `json:"Color"`
}

func main() {
	dog := Dog{"Bruce", 5, "White"}
	// serialize → json
	jsonDog, err := json.Marshal(dog)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonDog))

	nameTag := reflect.TypeOf(dog).Field(0).Tag.Get("json")
	ageTag := reflect.TypeOf(dog).Field(0).Tag.Get("json")
	colorTag := reflect.TypeOf(dog).Field(0).Tag.Get("json")
	fmt.Println(nameTag)
	fmt.Println(ageTag)
	fmt.Println(colorTag)
}
