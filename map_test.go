package slicehelpers

import "testing"

func TestMap(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5}
	strSlice := []string{"a", "b", "c"}
	floatSlice := []float64{1.1, 2.2, 3.3}

	mappedInt := Map(intSlice, func(v int) int {
		return v * 2
	})
	expectedInt := []int{2, 4, 6, 8, 10}

	for i, v := range mappedInt {
		if v != expectedInt[i] {
			t.Errorf("Expected %v, got %v", expectedInt, mappedInt)
		}
	}

	mappedStr := Map(strSlice, func(v string) string {
		return v + v
	})
	expectedStr := []string{"aa", "bb", "cc"}

	for i, v := range mappedStr {
		if v != expectedStr[i] {
			t.Errorf("Expected %v, got %v", expectedStr, mappedStr)
		}
	}

	mappedFloat := Map(floatSlice, func(v float64) float64 {
		return v * 2
	})
	expectedFloat := []float64{2.2, 4.4, 6.6}

	for i, v := range mappedFloat {
		if v != expectedFloat[i] {
			t.Errorf("Expected %v, got %v", expectedFloat, mappedFloat)
		}
	}
}