```go
package main

import "fmt"

func main() {

    s := make([]int, 0, 5)
    s = append(s, 1, 2, 3, 4, 5)

    // Check if the index is within the bounds of the slice
    if len(s) > 5 {
        fmt.Println(s[5])
    } else {
        fmt.Println("Index out of bounds")
    }

}
```