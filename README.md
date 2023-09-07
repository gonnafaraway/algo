# algo

## About
Library with various arithmetic functions for structured data.
Try it now!

## Usage
```go
package main

import (
	"github.com/gonnafaraway/algo/maps"
)

func main() {
	a := make(map[string]int)
	a["1"] = 1
	a["2"] = 4
	fmt.Println(maps.FindLargestValue(a)) // 4
}
```