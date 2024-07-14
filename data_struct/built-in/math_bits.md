Package bits implements bit counting and manipulation functions for the predeclared unsigned integer types.

Functions in this package may be implemented directly by the **compiler**, for **better performance**. 

UC: Encryption, Hash, Compression Alg...

```go
package main

import (
	"fmt"
	"math/bits"
)

func main() {
    // leading & trailing zeros
    fmt.Println(bits.LeadingZeros(8))    // 28 (32-bit)
	fmt.Println(bits.LeadingZeros(128))  // 24 (32-bit)
    fmt.Println(bits.TrailingZeros(8))   // 3  (32-bit)
	fmt.Println(bits.TrailingZeros(128)) // 7  (32-bit)
    
    // ones
    fmt.Println(bits.OnesCount(8))  // 1
	fmt.Println(bits.OnesCount(15)) // 4
    
    // rotate
    fmt.Printf("%08b\n", bits.RotateLeft(8, 1))   // 00010000
	fmt.Printf("%08b\n", bits.RotateRight(16, 1)) // 00001000
    
    // reverse bit & byte
    var x uint8 = 0b11010010
	reversed := bits.Reverse8(x)       // 0b01001011
    var y uint32 = 0x12345678
	reversed := bits.ReverseBytes32(y) // 0x78563412
    
    // valid length
    fmt.Println(bits.Len(8))   // 4
	fmt.Println(bits.Len(128)) // 8
    
    // 2, 0
    sum, carry := bits.Add(1, 1, 0)   // 2, 0
    diff, borrow := bits.Sub(2, 1, 0) // 1, 0
    
    sum, carry := bits.Add(0xFFFFFFFF, 1, 0) // 0, 1
    diff, borrow := bits.Sub(0, 1, 0)        // 0xFFFFFFFF, 1
}
```

