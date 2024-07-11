package slicehelpers

// Reduce reduces a slice to a single value by applying a function.
func Reduce[T any, U any](slice []T, reducer func(U, T) U, initial U) U {
	result := initial
	for _, v := range slice {
		result = reducer(result, v)
	}
	return result
}
