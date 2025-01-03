```go
package main

import "fmt"

func main() {

    // Slice with capacity 5 and length 0
    s := make([]int, 0, 5)

    // Appending elements to the slice
    s = append(s, 1, 2, 3, 4, 5)

    // Trying to access an element beyond the current length of slice
    fmt.Println(s[5]) // This will cause a runtime panic

}
```