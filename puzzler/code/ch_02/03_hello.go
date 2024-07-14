// *.go belongs to which package.
package main

// package to import, so we could use func inside package.
// $GOROOT/src/fmt
import "fmt"

// "main" func as entrypoint, run first if not init()
func main() {
	fmt.Println("Hello world")
}
