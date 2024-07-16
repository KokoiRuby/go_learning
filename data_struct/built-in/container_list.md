Package list implements a **doubly linked list**. 

:construction_worker: **​Not circular!** where head.prev & tail.next = nil.

:construction_worker: **Not thread-safe = sync primitives needed!**

```go
package main

import (
	"container/list"
	"fmt"
)

func main() {
    // new (delayed), malloc when used
	l := list.New()

	// add elem to back/front of the list
	l.PushBack("Element 1")
	l.PushFront("Element 0")

	// insert before/after a elem
	elem := l.PushBack("Element 2")
	l.InsertBefore("Element 1.5", elem)
	l.InsertAfter("Element 2.5", elem)

	// iter
	fmt.Println("List elements:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

    // get by iter only
    toFind := "Element 2"
    if elem := findElementByValue(l, toFind); elem != nil {
        // TODO
    } else {
        // TODO
    }
    
    // move
    e1 := l.PushBack("Element 1")
	e2 := l.PushBack("Element 2")
	e3 := l.PushBack("Element 3")
	e4 := l.PushBack("Element 4")
    
    l.MoveAfter(e1, e3)
    l.MoveBefore(e4, e2)
    l.MoveToBack(e2)
    l.MoveToFront(e3)
    
	// remove a elem
	l.Remove(elem)

	// get len of list = # of elem in list
	fmt.Println("Length:", l.Len())

	// get front/back elem
	fmt.Println("Front:", l.Front().Value)
	fmt.Println("Back:", l.Back().Value)
    
    // init or clean up
    l.Init()
}

func findElementByValue(l *list.List, value interface{}) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return e
		}
	}
	return nil
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
```

