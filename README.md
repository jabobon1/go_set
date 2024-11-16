# Go Generic Set Implementation

Implementation of Set structure in GO

## Installation

```bash
go get github.com/jabobon1/go_set
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/jabobon1/go_set/pkg/set"
)

func main() {
    // Create a new integer set
    numbers := set.New[int](10)
    
    // Add elements
    numbers.Add(1).Add(2).Add(3)
    
    // Check membership
    fmt.Println(numbers.In(2))  // true
    fmt.Println(numbers.In(4))  // false
    
    // Get all elements
    fmt.Println(numbers.All())  // [1 2 3]
}
```

## Usage Examples

### Basic Operations

```go
// Create a new set
strSet := set.New[string](5)

// Add elements (chainable)
strSet.Add("apple").Add("banana").Add("cherry")

// Check if element exists
exists := strSet.In("apple")    // true

// Get number of elements
size := strSet.Len()            // 3

// Remove an element
strSet.Remove("banana")

// Get all elements
elements := strSet.All()        // ["apple", "cherry"]

// Clear the set
strSet.Clear()
```

### Set Operations

```go
set1 := set.New[int](5)
set2 := set.New[int](5)

set1.Add(1).Add(2).Add(3)
set2.Add(2).Add(3).Add(4)

// Union
union := set1.Union(*set2)          // [1 2 3 4]

// Intersection
intersection := set1.Intersect(*set2) // [2 3]

// Symmetric Difference
diff := set1.Diff(*set2)             // [1 4]
```

## Running Tests

```bash
# Run all tests
make test

# Run benchmarks
make bench

# Check test coverage
make cover
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by Go's internal map implementation
- Built with modern Go generics support
- Designed with performance and usability in mind

## Project Status

This project is actively maintained. Feature requests and bug reports are welcome.

## Author

Your Name ([@jabobon](https://github.com/jabobon1))