package maps

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
