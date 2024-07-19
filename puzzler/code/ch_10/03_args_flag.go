package main

import (
	"flag"
	"fmt"
)

func main() {
	stringFlag := flag.String("name", "world", "a name to greet")
	intFlag := flag.Int("age", 30, "an age to display")
	boolFlag := flag.Bool("verbose", false, "enable verbose output")

	flag.Parse()

	fmt.Printf("Hello, %s!\n", *stringFlag)
	fmt.Printf("You are %d years old.\n", *intFlag)
	if *boolFlag {
		fmt.Println("Verbose mode enabled.")
	} else {
		fmt.Println("Verbose mode disabled.")
	}
}
