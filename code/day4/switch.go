package main

import "fmt"

func main() {
	var score int
	var grade string
	fmt.Println("Input a score: ")
	fmt.Scanf("%d", &score)

	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		// grade = "D"
		fallthrough
	default:
		grade = "F"
	}

	fmt.Println("Grade is: ", grade)
}
