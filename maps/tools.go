package maps

import "fmt"

// FindLargestValue helps to find the largest value in a maps.
func FindLargestValue[T comparable](m map[T]int) int {
	var largestValue int

	for _, v := range m {
		if v > largestValue {
			largestValue = v
		}
	}

	return largestValue
}

// Print prints to console any type of map.
func Print[T comparable](m map[T]T) {
	for k, v := range m {
		fmt.Printf("%#v: %#v\n", k, v)
	}
}
