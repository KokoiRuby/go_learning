package main

import (
	"custom/pkg1"
	"fmt"
	// "github.com/gin-gonic/gin"
)

func main() {
	var test1 string
	test1 = pkg1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pkg1.Pack1Int)
	fmt.Printf("Float from package1: %f\n", pkg1.PackFloat)
}
