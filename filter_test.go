package slicehelpers

import "testing"

func TestFilter(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	strSlice := []string{"a", "b", "c"}
	floatSlice := []float64{1.1, 2.2, 3.3}

	filteredInt := Filter(intSlice, func(v int) bool {
		return v%2 == 0
	})
	expectedInt := []int{2, 4}

	for i, v := range filteredInt {
		if v != expectedInt[i] {
			t.Errorf("Expected %v, got %v", expectedInt, filteredInt)
		}
	}

	filteredStr := Filter(strSlice, func(v string) bool {
		return v != "b"
	})
	expectedStr := []string{"a", "c"}

	for i, v := range filteredStr {
		if v != expectedStr[i] {
			t.Errorf("Expected %v, got %v", expectedStr, filteredStr)
		}
	}

	filteredFloat := Filter(floatSlice, func(v float64) bool {
		return v > 2.0
	})
	expectedFloat := []float64{2.2, 3.3}

	for i, v := range filteredFloat {
		if v != expectedFloat[i] {
			t.Errorf("Expected %v, got %v", expectedFloat, filteredFloat)
		}
	}
}
