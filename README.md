# Go Slice Out of Bounds Error

This repository demonstrates a common error in Go programming: accessing elements beyond the current length of a slice.  Attempting to access an index outside the valid range of a slice will result in a runtime panic.

## The Bug

The `bug.go` file contains code that appends elements to a slice.  It then attempts to access an element beyond the current length of the slice.  This will trigger a runtime panic.

## The Solution

The `bugSolution.go` file demonstrates how to correctly handle slice access by checking the index against the slice's length before accessing any elements.

## How to Run

1. Clone the repository.
2. Navigate to the repository's directory.
3. Run the `bug.go` file. Observe the panic.
4. Run the `bugSolution.go` file. Observe the corrected output without panic.
