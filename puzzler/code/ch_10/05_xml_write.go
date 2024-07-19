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
	person := Person{
		Name:  "Jane Doe",
		Age:   28,
		Email: "janedoe@example.com",
	}

	file, err := os.Create("to_write.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// xml encoder
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")

	// write header
	_, err = file.WriteString(xml.Header)
	if err != nil {
		fmt.Println("Error writing XML header:", err)
		return
	}

	// encode
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding XML:", err)
		return
	}

	fmt.Println("XML data written to to_write.xml")
}
