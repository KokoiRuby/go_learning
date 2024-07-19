package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func writeGobToFile(filename string, person Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		return err
	}

	return nil
}

func readGobFromFile(filename string) (Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Person{}, err
	}
	defer file.Close()

	var person Person
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&person)
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

	err := writeGobToFile("person.gob", originalPerson)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	decodedPerson, err := readGobFromFile("person.gob")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Printf("Decoded Person from file: %+v\n", decodedPerson)
}
