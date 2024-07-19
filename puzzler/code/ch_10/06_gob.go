package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func encodeGob(person Person) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(person)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func decodeGob(data []byte) (Person, error) {
	var person Person
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(&person)
	if err != nil {
		return person, err
	}
	return person, nil
}

func main() {
	originalPerson := Person{
		Name:  "Jane Doe",
		Age:   28,
		Email: "janedoe@example.com",
	}

	encodedData, err := encodeGob(originalPerson)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	fmt.Println("Encoded Data:", encodedData)

	decodedPerson, err := decodeGob(encodedData)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Printf("Decoded Person: %+v\n", decodedPerson)
}
