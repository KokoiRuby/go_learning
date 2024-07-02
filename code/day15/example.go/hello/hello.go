package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Alice", "Bob", "Cathy"}

	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)
	// message, err := greetings.Hello("Gladys")
	// message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
