[golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go.

:construction_worker: **Not thread-safe = sync primitives needed!**

What is considered **comparable** in Go?

- `Booleans`, `integers`, `strings`, `floats` or basically primitive types.
- `Pointers`
- `Arrays`
- `Structs` if *all of their fields* are also comparable independently.

```go
package main

import (
    "fmt"
    mapset "github.com/deckarep/golang-set/v2"
)

func main() {
    // Create a string-based set of required classes.
    required := mapset.NewSet[string]()
    required.Add("cooking")
    required.Add("english")
    required.Add("math")
    required.Add("biology")

    // Create a string-based set of science classes.
    sciences := mapset.NewSet[string]()
    sciences.Add("biology")
    sciences.Add("chemistry")
  
    // Create a string-based set of electives.
    electives := mapset.NewSet[string]()
    electives.Add("welding")
    electives.Add("music")
    electives.Add("automotive")

    // Create a string-based set of bonus programming classes.
    bonus := mapset.NewSet[string]()
    bonus.Add("beginner go")
    bonus.Add("python for dummies")
    
    // union, diff, intersect
    all := required
      .Union(sciences)
      .Union(electives)
      .Union(bonus)
    fmt.Println(all)
    notScience := all.Difference(sciences)
    reqScience := sciences.Intersect(required)
  
    // if had
    result := sciences.Contains("cooking")
    
    // distinct elem
    fmt.Println(bonus.Cardinality())
}
```

