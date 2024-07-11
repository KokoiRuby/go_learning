Package ring implements operations on **circular lists**.

:construction_worker: **Not thread-safe = sync primitives needed!**

```go
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// new
	r := ring.New(5)

	// iter
	fmt.Println("Iterate ring:")
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})

	// add elem
	r.Value = 1
	r = r.Next()
	r.Value = 2
	r = r.Next()
	r.Value = 3
	r = r.Next()
	r.Value = 4
	r = r.Next()
	r.Value = 5

	// iter
	fmt.Println("\nIterate ring:")
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})

	// links a new node where r.Next() becomes newNode
	newNode := ring.New(1)
	newNode.Value = 6
	r.Link(newNode)

	// iter
	fmt.Println("\nIterate ring:")
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})

	// move pointer forward 3 steps
	r = r.Move(3)
    // move pointer backward 3 steps
	// r = r.Move(-3)
    
    // iter
	fmt.Println("\nIterate ring:")
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
    
	r.Unlink(2)   // 删除后面的 2 个节点

	// iter
	fmt.Println("\nIterate ring:")
	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
}
```

