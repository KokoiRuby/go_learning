Package sync provides **basic sync primitives** such as mutex locks. Other than the Once and WaitGroup types, most are intended for use by **low-level** library routines. **Higher-level** synchronization is better done via **channels** and communication.

Map is like a Go map[any]any but is **safe for concurrent use by multiple goroutines** without additional locking or coordination. Loads, stores, and deletes run in amortized constant time.

:smile:

1. When the entry for a given key is only ever written once but **read many times**, as in caches that only grow.
2. When multiple goroutines read, write, and overwrite entries for **disjoint** sets of keys.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var sm sync.Map

    // store
    sm.Store("key1", "value1")
    sm.Store("key2", 42)

    // load
    if value, ok := sm.Load("key1"); ok {
        fmt.Println("Loaded key1:", value)
    } else {
        fmt.Println("Key1 not found")
    }

    // store if not found, load if found
    if value, loaded := sm.LoadOrStore("key3", "value3"); loaded {
        fmt.Println("Loaded existing key3:", value)
    } else {
        fmt.Println("Stored new key3 value")
    }

    // delete
    sm.Delete("key1")

    // iter
    sm.Range(func(key, value interface{}) bool {
        fmt.Println("Key:", key, "Value:", value)
        return true
    })
    
    // compare & deleted if same val
    sm.Store("key1", "value1")
    deleted := sm.CompareAndDelete("key1", "value1")
    
    // compare & swap if same val
    sm.Store("key1", "value1")
    swapped := sm.CompareAndSwap("key1", "value1", "value2")
    
    // load & deleted
    sm.Store("key1", "value1")
    value, loaded := sm.LoadAndDelete("key1")
    
    // swap & return prev
    sm.Store("key1", "value1")
    previous, loaded := sm.Swap("key1", "value2")
    
}

```

