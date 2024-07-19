package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func main() {
	file, err := os.Open("to_read.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// xml decoder
	decoder := xml.NewDecoder(file)

	// struct to keep xml doc
	var person Person

	// decode
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Email: %s\n", person.Email)
}
