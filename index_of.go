package slicehelpers

// IndexOf returns the index of the first occurrence of an element in a slice, or -1 if not found.
func IndexOf[T comparable](slice []T, element T) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}
