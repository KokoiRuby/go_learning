package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%v", now)
	fmt.Printf("%T", now)

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	t := time.Date(1992, time.December, 7, 12, 0, 0, 0, time.UTC)
	fmt.Println(t)

}
